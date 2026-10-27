# Hyperliquid (HyperCore) Protobuf Schemas

Canonical protobuf schemas for Bitquery's Hyperliquid **HyperCore** streaming pipeline.

HyperCore is the native execution engine of the Hyperliquid L1 — the on-chain perpetual
and spot exchange (order books, matching, liquidations, funding, staking, vaults). It is
distinct from **HyperEVM**, the EVM side of the same chain; the two are linked, and one of
the raw streams below (`CoreWriterActions`) records exactly that link.

The package has **two layers**:

* **Raw ingestion** (`hypercore.proto`, `transactions.proto`, and the per-stream files).
  A **lossless, typed re-encoding** of the NDJSON the official Hyperliquid node
  (`hl-visor run-non-validator`) writes under `~/hl/data`: nothing is aggregated, filtered
  or re-interpreted. Where the node's JSON uses tagged unions, the schema flattens them
  into a tag field plus optional variant fields; where a field is unknown or unstable it is
  kept verbatim in an `Extra` map or as raw JSON bytes, so no information is ever dropped.
  This layer is **internal** — it is how block data enters the pipeline.
* **The delivered product** (`hyperliquid_block_message.proto`, `candles.proto`). A
  **projection** of the raw layer into one consumer-facing message per block:
  `HyperliquidCoreBlockMessage` (all trading activity) and `CandlesBlockMessage` (OHLC).
  **This is what subscribers receive.** It is a normalization, not a re-encoding: the raw
  streams are reassembled, decoded into typed sub-objects, and sorted into per-kind lists.
  Next to it, `hyperliquid_snapshot_message.proto` (`HyperliquidSnapshotMessage`) delivers
  the exchange **state** every 10,000 blocks — positions, margin, the L4 book with owners,
  spot balances, Earn, vaults — projected from the node's own state snapshot, not from the event
  streams (§7).

If you consume the Bitquery stream, start at **§5 (the product)** and **§7 (snapshots)**.
§2–§4 (conventions, identity, symbology) underlie both layers. §8 documents the raw node
streams for reference.

**Schema evolution.** When the node starts emitting a new field, it first appears in
`Extra` (§2.4) on the raw layer — already usable, as a raw JSON string. The upgrade path
is: add a typed `optional` field under a **fresh** field number, and from that release on
the value arrives typed. Historical messages are **not rewritten** — they keep the value
in `Extra` under the node's original field name. Proto field numbers are never reused or
renumbered, so old and new binaries stay compatible both ways.

> **Provenance.** Everything in this document is verified against one of:
> (1) the `.proto` files in this directory, (2) real mainnet node output (samples and
> multi-million-event audits, July 2026), or (3) official Hyperliquid documentation —
> [node README](https://github.com/hyperliquid-dex/node),
> [L1 data schemas](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/nodes/l1-data-schemas),
> [historical data](https://hyperliquid.gitbook.io/hyperliquid-docs/historical-data).
> Facts that come from measurement rather than specification are marked *observed*.

---

## 1. The big picture

```
Hyperliquid node (hl-visor) writes NDJSON under ~/hl/data
 ├── 7 event streams   (fills, order_statuses, raw_book_diffs, twap_statuses,
 │                       hip3_oracle_updates, misc_events, core_writer_actions)
 ├── replica_cmds       (consensus transactions — what went INTO each block)
 └── periodic_abci_states (msgpack state dump every 10,000 blocks — archived as is, §7)
        │
        ▼  lossless typed re-encoding   (raw ingestion — internal)
 ├── HyperCoreBlock / HyperCoreBlocks              (hypercore.proto)       — the 7 event streams
 └── HyperCoreTransactionsBlock / …Blocks          (transactions.proto)    — consensus transactions
        │
        ▼  projection / normalization   (the product)
 ├── HyperliquidCoreBlockMessage                   (hyperliquid_block_message.proto)    ──▶ subscribers
 ├── CandlesBlockMessage                           (candles.proto)                      ──▶ subscribers
 └── HyperliquidSnapshotMessage  ◀── the .rmp      (hyperliquid_snapshot_message.proto) ──▶ subscribers
```

| Message | Role | Contents |
|---|---|---|
| **`HyperliquidCoreBlockMessage`** | **delivered** — one per block | All trading activity of a block, one list per kind: trades, order updates, funding, liquidations, TWAP lifecycle, book deltas, prices, leverage changes, signed consensus actions |
| **`CandlesBlockMessage`** | **delivered** — one per trades block | OHLC bars touched by the block (1m / 5m / 15m / 1h) |
| **`HyperliquidSnapshotMessage`** | **delivered** — one snapshot per 10,000 blocks, in parts | Exchange state at one block: markets, accounts per perp dex, open positions, every book and trigger order with its owner, spot balances, vaults and their depositors |
| `HyperCoreBlock` | raw ingestion — internal | The 7 node event streams for one block, re-encoded losslessly |
| `HyperCoreTransactionsBlock` | raw ingestion — internal | The consensus `replica_cmds` for one block: every signed action bundle with signatures, nonces and the engine's accept/reject verdict |

The product answers "what happened, and to whom?". The raw transactions layer additionally
answers "who asked for it, and was it accepted?" — rejections that never produce events
(bad nonce, malformed action, insufficient margin at the API level) survive only there, and
the product surfaces them through its `SignedActions` list (§5.9).

### Blocks

* HyperCore has **no block hash**. `BlockNumber` is the canonical block identifier.
* `BlockNumber ↔ BlockTime` is a bijection — no two blocks share a timestamp
  (*observed*: 374,720 distinct pairs in 374,720 blocks over 12 h).
* `BlockTime` is the consensus timestamp in **epoch nanoseconds (UTC)**, identical across
  every stream — and every product list — of the same block.
* Blocks are fast: roughly 10–15 blocks/s on mainnet (*observed*, July 2026 — block
  numbers were around 1.07–1.08 billion).
* A block in which a given kind of event did not occur carries an **empty list** for it —
  that is the shape, not a gap.
* `Round` / `ParentRound` (consensus counters) are **not** block numbers: empty rounds
  produce no block, so `Round` runs ahead of `BlockNumber` (~28% ahead, *observed*).

### File / message layout (raw archive)

One protobuf message of a container type packs many blocks:

* `HyperCoreBlocks.Blocks` — repeated `HyperCoreBlock`
* `HyperCoreTransactionsBlocks.Blocks` — repeated `HyperCoreTransactionsBlock`

In Bitquery's archive each stored object covers exactly one 1000-block bucket
(`[K*1000 … K*1000+999]`).

---

## 2. Conventions (read this before anything else)

### 2.1 Numbers are decimal strings

Every monetary value — price, size, fee, PnL, notional, funding — is a **decimal string**
(`"7541.5"`, `"0.002606"`), exactly as the node prints it. Never floats: floats lose
precision on financial data. Parse with a decimal type.

Two deliberate exceptions:

* `core_writer_actions.proto`, where `Token`, `Wei` and `Ntl` are raw integers (a token
  index, a wei amount, an integer notional), because that is what the node emits there.
* `candles.proto`, where OHLC values are `double` (Float64) — a candle is already an
  aggregate, and Float64 keeps full precision even for high-priced assets (BTC ~100k).

### 2.2 Timestamps are mixed — per field, not per stream

The node mixes epoch-millisecond integers and ISO nanosecond strings. The schemas keep
each field's native precision (ISO strings are parsed to integer nanoseconds). The product
preserves the source precision of each field it carries.

| Field | Unit | Meaning |
|---|---|---|
| `*.BlockTime` (every container / product) | **ns** | consensus block time |
| `Fill.Time`, `Execution.EventTime` | **ms** | fill event time |
| `OrderStatus.Time`, `OrderUpdate.EventTime` | **ns** | order status event time |
| `Order.Timestamp`, `OrderUpdate.PlacedTime` | **ms** | order placement time |
| `TwapStatus.Time`, `TwapInterval.EventTime` | **ns** | TWAP event time |
| `TwapState.Timestamp`, `TwapInterval.StartTime` | **ms** | TWAP order **start** time (not event time) |
| `MiscEvent.Time` | **ns** | event time |
| `OraclePx.LastUpdateTime`, `PriceUpdate.UpdateTime` | **ns** | last oracle update |
| `SignedAction.Nonce`, `SignedActionEvent.Nonce`, `Trader.SignedAt` | ms *(by wallet convention)* | signer nonce / signing time |
| `SignedAction.ExpiresAfter`, `SignedActionEvent.ExpiresAfter` | ms | action deadline, `0` = absent |
| `TimeInterval.Duration/Start/End` (candles) | **s** | candle bucket boundaries |
| `MarketPrices.UpdateTime` (snapshot) | **ns** | last oracle update |
| `OrderState.PlacedTime` (snapshot) | **ms** | order placement time (same as `OrderUpdate.PlacedTime`) |
| `VaultState.CreateTime`, `VaultFollowTimes.EntryTime` / `LastDepositTime` / `LockupUntil` (snapshot) | **ns** | vault creation, first and last deposit, end of lockup |
| `EarnReserveState.UpdateTime` (snapshot) | **ns** | the last interest update of an Earn token |
| `VaultTerms.LockupSeconds` (snapshot) | **s** | lockup after each deposit |

### 2.3 Addresses and hashes are `bytes` (raw) or `0x…` strings (product)

On the **raw** layer, `0x…` hex strings from the node are decoded to raw `bytes`: 20 bytes
for addresses, 32 bytes for hashes; empty/absent → empty bytes or unset `optional`. On the
**product** layer the trading messages (`Trader`, `Market`, `Execution`, …) carry addresses
and hashes as human-readable **`0x…` strings**; the consensus-sourced `SignedActionEvent`
keeps them as `bytes` (it is a thin denormalization of the raw consensus row).

### 2.4 `Extra` maps — the forward-compatibility net (raw layer)

Nearly every **raw** message carries `map<string,string> Extra`. Any JSON key the node
emits that the schema does not model lands there — key = the node's field name, value = the
raw JSON value. In the common case `Extra` is empty; a non-empty `Extra` means the node
started emitting a new field and nothing was lost. The product does not carry `Extra`
(it is a curated projection), except that `SignedActionEvent` keeps the raw action and
response JSON verbatim in `Action` / `Response`.

### 2.5 Flattened tagged unions

The node encodes variants as `{"tagName": {...}}` objects or bare strings. The schemas
flatten these into a string tag field (`Kind`, `Status`, `Type`, `InnerType`,
`ActionType`) set **verbatim**, plus per-variant fields. An unseen variant still comes
through: the tag is preserved and its payload goes to `Extra`.

### 2.6 `LowCardinality`

Comments marked `LowCardinality` are hints for columnar storage (e.g. ClickHouse
`LowCardinality(String)`): the field draws from a small set of values, even when that set
is open.

---

## 3. Identity: how things are keyed and joined

HyperCore has **three distinct levels** — do not conflate them:

```
TRANSACTION (Hash)  →  MATCH (Coin, Tid)  →  SIDE (User)
```

* One L1 transaction (`Hash`) can produce **many** matches — a single aggressive order
  sweeping the book was *observed* to generate 482 distinct `Tid`s across 80 coins under
  one `(BlockNumber, Hash)`. `Hash` identifies an *action*, not a trade.
* One match normally produces **two** fills — one per counterparty. In the product these
  are two `Trade` rows (or `PerpLiquidation` rows); pair them on `(BlockNumber, Market, Tid)`.

| Identifier | What it is | Unique? | Safe usage |
|---|---|---|---|
| `BlockNumber` | canonical block id | globally unique | primary block key |
| `Oid` | exchange-assigned order id (uint64) | unique per order | order lifecycle joins |
| `Cloid` | optional client-supplied order id (16 bytes) | client-controlled | client-side matching only |
| `Tid` | **50-bit** hash of `(buyer_oid, seller_oid)` | **no** — collides at scale (~2²⁵ trades) | unique only within one block+coin |
| `Hash` | L1 transaction hash | **no** — shared by many fills; **all-zero on both sides of TWAP fills** | display / grouping, never dedup |
| `TwapId` | TWAP order id | unique per TWAP | see caveat in §5.5 |
| **`(BlockNumber, Coin/Market, Tid)`** | **the trade key** | **yes** (*observed* on millions of fills) | dedup, pairing the two sides of a match |
| `Execution.Index` | place of the fill in the block's fill list, from 1 | unique within a block | **the only order of fills inside a block** |

**`Tid` is not an order.** Being a hash, it says nothing about which fill came first: inside
one block, ordering fills by `Tid` picks an arbitrary one as "the last". A position size, a
candle's open/close or an average entry price taken that way is wrong whenever an account or a
market has several fills in the block. Order by `(BlockNumber, Execution.Index)`; the node
lists fills in execution order, so `Index` is that order, and `SizeBefore` of each fill equals
`Size` of the previous fill of the same account and coin.

**Event hashes are synthetic.** `Fill.Hash` / `OrderStatus.Hash` / `MiscEvent.Hash`
(and thus `Execution.Hash` in the product) is *not* a hash of the transaction content.
Starting at byte 10 it embeds a compact length-prefixed big-endian encoding of the event's
position:

```
[10 bytes][0x04][u32 block number][len][action seq no][len][sub-action idx][~13 bytes]
```

so the hash itself decodes to `(BlockNumber, action-sequence-in-block, sub-order index)`
(*observed*: all 1,406 unique non-zero hashes of a full mainnet block plus samples from
other blocks and streams follow this layout, zero exceptions; July 2026). The remaining
23 bytes are engine-generated; they were verified **not** to equal the SDK `action_hash`
(keccak of msgpack action + nonce + vault), nor keccak of the action JSON, signature, or
bundle hash. The embedded action sequence follows the engine's execution order, which
does not match the bundle order recorded in `replica_cmds`, so do not use it to index
into bundles directly.

This synthetic hash is exactly the **official transaction hash**: the public explorer
API (`POST https://rpc.hyperliquid.xyz/explorer`, `{"type":"txDetails","hash":…}`)
resolves an event hash to its transaction — returning the same block number the hash
embeds, the acting user and the full action (*verified live: a fill hash from the node
stream resolved to the taker's IOC order in the same block, and the public
`userFillsByTime` info endpoint returned the same fill with the same hash*).

**Joining trades to the signed action.** The consensus bundle `Hash`
(`SignedActionBundle.Hash`) lives in a different universe than event hashes
(*verified disjoint*). The precise join is three-step (*verified on a full mainnet block*):

1. `Execution.Hash` equals the `Hash` of the aggressor order's `open`/`filled`
   `OrderUpdate` in the same block — both sides of a match carry the **taker's**
   transaction hash (34/34 fills in the verified block).
2. That `OrderUpdate` carries the `Oid`.
3. The `Oid` appears in the engine response (`SignedActionEvent.Response`, from the node's
   `resps`) of exactly one signed action (1,424/1,437 `open` statuses in the verified
   block) — giving the precise bundle and action.

TWAP fills (all-zero hash) attribute via the trade-key triple instead (§5.5). Where the
oid route is unavailable (e.g. rejected actions that never produced an order id), fall
back to the coarse join `BlockNumber` + `User`.

---

## 4. Market (coin) symbology

On the raw layer the `Coin` field is a raw market identifier; on the product layer it is
`Market.CoinRaw` (with `Market.Symbol`, `Market.Kind` and `Market.IsPerp` derived from it,
see §5.2). Four families:

| Pattern | Family | Example | Notes |
|---|---|---|---|
| plain symbol | native perpetual | `BTC`, `HYPE`, `kBONK` | `k`-prefix = 1000× denomination (e.g. `kBONK` = 1000 BONK) |
| `<dex>:<SYM>` | HIP-3 (builder-deployed) perpetual | `xyz:SP500`, `flx:BTC`, `mkts:AAPL` | prefix = the deployer's dex namespace; equities/FX/commodities live here |
| `@<n>` | spot pair by index | `@107` | n = spot pair index in exchange metadata |
| `#<n>` | outcome token (prediction market) | `#1890` | comes in complementary pairs `(n, n^1)` — see below |

These map to `Market.Kind` as `perp` / `hip3` / `spot` / `outcome`; `Market.IsPerp` is true
for the first two (§5.2).

**Outcome tokens** (`#<n>`): the two complementary outcomes of one market are the pair
`(n, n XOR 1)` — `#1890`/`#1891`. A *complete-set* operation (mint/burn of both outcomes at
once; `Dir` values like `Buy`+`Buy`, `Sell`+`Sell`, `Split Outcome`, `Merge Outcome`,
`Merge Question`) emits **one fill per leg**: same `Hash`, same `Sz`, same `Dir`, but
**different `Coin` and `Tid`**, with the two `Px` values summing to exactly 1. Fee is
`0.0` and `FeeToken` is the outcome token itself, written `+<n>`. (*Observed* on 10,125
such pairs.) This is why some matches look one-sided under the trade-key triple: the
counterpart leg lives under the sibling coin — it is not lost. A trades pipeline must
therefore **never** inner-join the two sides of a match.

**Fee-token notation** `+<n>`: a token referenced by numeric index (seen for outcome
tokens and some spot tokens, e.g. `+1891`, `+7970`).

---

## 5. The delivered product — `HyperliquidCoreBlockMessage`

`hyperliquid_block_message.proto`. **One message per HyperCore block** — the single-stream
canon a subscriber receives. Every kind of event a block produced rides in its own list, so
one message carries the whole block and a downstream recorder projects each list into its
own table.

**All market kinds flow here** — perp, HIP-3, spot and outcome — told apart by `Market.Kind`
and `Market.IsPerp`. Perp markets (perp + hip3) carry the full object set
(Position/Leverage + Funding/Liquidation events); spot and outcome share the same execution
core but leave the perp-only fields empty. **Filter clean perps on `Market.IsPerp`.**

### 5.1 Envelope

| Field | Type | Description |
|---|---|---|
| `BlockNumber` | uint64 | canonical block id |
| `BlockTime` | int64 | consensus block time, epoch **ns** |
| `Trades` | repeated Trade | executions, non-liquidation (§5.3) |
| `Orders` | repeated OrderUpdate | order lifecycle: placements, fills, cancels, rejects, TP/SL legs (§5.4) |
| `Funding` | repeated PerpFunding | hourly funding, per (trader, market) (§5.7) |
| `Liquidations` | repeated PerpLiquidation | forced closes — their own product, excluded from `Trades` (§5.3) |
| `Twaps` | repeated Twap | TWAP lifecycle (§5.5) |
| `BookUpdates` | repeated BookUpdate | per-order (L4) book changes (§5.6) |
| `Prices` | repeated PriceUpdate | oracle / mark price republishes (§5.8) |
| `LeverageChanges` | repeated TraderLeverageUpdate | accepted `updateLeverage` — from the consensus stream (§5.7) |
| `SignedActions` | repeated SignedActionEvent | consensus actions: rejects + rare types — from the consensus stream (§5.9) |

An **empty list** means the block had nothing of that kind. Seven lists are sourced from
the core event streams; **`LeverageChanges` and `SignedActions` come from the consensus
stream** (`transactions.proto`).

### 5.2 Shared sub-messages

Reused across the lists above.

**`Trader`** — the acting parties (all `0x…` strings):

| Field | Type | Description |
|---|---|---|
| `Address` | string | master account |
| `Signer` | string | signing agent (recovered via ecrecover) |
| `Vault` | string | vault address, when acting for one |
| `Broadcaster` | string | node/relay that broadcast the action |
| `SignedAt` | int64 | signing time, epoch **ms** |

**`Market`** — the resolved market (from `CoinRaw`, §4):

| Field | Type | Description |
|---|---|---|
| `Symbol` | string | display symbol |
| `Kind` | string | `"perp"` \| `"hip3"` \| `"spot"` \| `"outcome"` |
| `Protocol` | string | `"hypercore"` (native), else the HIP-3 dex namespace |
| `CoinRaw` | string | raw coin id: `"HYPE"`, `"xyz:NFLX"`, `"@107"`, `"#1890"` |
| `IsPerp` | bool | `true` for perp + hip3 (full object set); `false` for spot / outcome — **filter on this** |
| `MaxLeverage` | uint32 | max leverage for the perp market (info-API meta); `0` = unknown / spot / outcome |

**`PerpPosition`** — the trader's position **as of** the event. Only realtime fields are
present; the snapshot fields (collateral, liquidation price, unrealized PnL) are **absent**
because HL exposes them only via a snapshot/API, not the realtime feed. **Empty = source
silent, NOT zero.**

| Field | Type | Description |
|---|---|---|
| `Side` | string | `"Long"` \| `"Short"` |
| `Size` | string | signed size **after** this execution (`Trade`, `PerpLiquidation`); for spot, the base-token balance net of a fee charged in that token; empty on order and book events |
| `SizeBefore` | string | size before this event |
| `EntryPrice` | string | empty if the position was opened before the window |
| `RealizedPnl` | string | realized PnL |
| `Funding` | string | signed |
| `Leverage` | uint32 | `0` = not yet in the leverage map (**NOT** 1×) |
| `IsCross` | bool | cross vs isolated |
| `Closed` | bool | position closed by this event |

**`Fees`** — money charged/credited on an execution (reused by `Trade` and
`PerpLiquidation`; only executions carry fees). `BuilderFee` and `DeployerFee` are
**components** of `Fee` (same unit); `PriorityGas` is **not**.

| Field | Type | Description |
|---|---|---|
| `Fee` | string | negative = maker rebate |
| `FeeToken` | string | not always USDC (§5.3) |
| `BuilderFee` | string | portion of `Fee` to the builder |
| `DeployerFee` | string | portion of `Fee` to the HIP-3 asset deployer |
| `PriorityGas` | string | **NOT** a component of `Fee`; may be in a different unit |

**`Execution`** — the execution itself, the shared core of `Trade` and `PerpLiquidation`
(both come from one `Fill`). Pair both sides of a match on `(BlockNumber, Market, Tid)`;
order the fills of a block by `Index`, never by `Tid` (§3).

| Field | Type | Description |
|---|---|---|
| `Side` | string | `"Buy"` \| `"Sell"` |
| `Direction` | string | `"Open Long"`, `"Close Short"`, … (open set) |
| `IsAggressor` | bool | `true` = taker |
| `Price` | string | execution price |
| `Size` | string | executed size |
| `Oid` | uint64 | order id |
| `Cloid` | string | client order id |
| `Tid` | uint64 | match id — join both sides on `(BlockNumber, Market, Tid)`; a hash, **not an order** (§3) |
| `Hash` | string | L1 tx hash; **zero on TWAP** — never a join key (§3) |
| `EventTime` | int64 | per-fill event time, epoch **ms** (finer than block `BlockTime`) |
| `Index` | uint32 | place of the fill in the block's fill list, from 1 — the execution order; one sequence for trades and liquidations; `0` on data written before the field existed |

### 5.3 Trades & liquidations (`Trade`, `PerpLiquidation`)

Both come from a node `Fill`. A **non-liquidation** fill becomes a `Trade`; a fill with the
liquidation flag set is routed to `PerpLiquidation` **only** (excluded from `Trades`), so
that row is the sole record — nothing is lost. A liquidation *is* an execution, so it shares
the same `Execution` / `Fees` / `PerpPosition` objects plus four liquidation-only fields.

**`Trade`:**

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | acting parties |
| `Market` | Market | resolved market |
| `Execution` | Execution | side / price / size / oid / tid / hash / … |
| `Fees` | Fees | fee breakdown |
| `Position` | PerpPosition | position as of the fill |
| `MarkPx` | string | mark price; empty if no oracle republish this block |
| `Builder` | string | builder address (usually empty) |
| `IsTwap` | bool | this fill is part of a TWAP |
| `TwapId` | uint64 | `0` when not a TWAP |

**`PerpLiquidation`** — `Trader` / `Market` / `Execution` / `Fees` / `Position` as above,
plus the four fields that make it a liquidation:

| Field | Type | Description |
|---|---|---|
| `MarkPx` | string | mark price **at liquidation** — NOT the execution price (`Execution.Price`) |
| `Method` | string | `"market"` \| `"backstop"` (backstop = HLP absorbed it — a market-stress signal) |
| `Liquidator` | string | backstop liquidator address, when the source names one |
| `LiquidatedUser` | string | `0x…` address being liquidated |

**Fee-token & fee algebra** (from the underlying `Fill`): `FeeToken` is an open set, not
always `"USDC"` — spot fills may pay in the spot token (`HYPE`, `PURR`, `UBTC`, …), outcome
markets in `+<n>` (~3.7% of fills are non-USDC, *observed*). `BuilderFee` and `DeployerFee`
are already **included** in `Fee`; `PriorityGas` is **separate** — do not add it to `Fee`.

### 5.4 Order updates (`OrderUpdate`)

One order-lifecycle event — the complete L1 order audit trail, including orders that never
reached the book (rejects). Sourced from the node `OrderStatus` stream.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | order owner |
| `Market` | Market | resolved market |
| `Side` | string | `"Buy"` \| `"Sell"` |
| `Status` | string | `"filled"` \| `"open"` \| `"canceled"` \| `"badAloPxRejected"` \| … (open set) |
| `IsReject` | bool | true for the many `*Rejected` statuses |
| `OrderType` | string | `"Limit"` \| `"Market"` \| `"Stop Market"` \| … |
| `Tif` | string | `"Alo"` \| `"Ioc"` \| `"Gtc"` \| … |
| `ReduceOnly` | bool | order may only reduce a position |
| `LimitPx` | string | limit price |
| `Size` | string | **remaining** size |
| `OrigSz` | string | original size |
| `TriggerPx` | string | trigger price |
| `TriggerCondition` | string | e.g. `"Price above 917.08"` |
| `IsTrigger` | bool | stop / take-profit order |
| `IsPositionTpsl` | bool | TP/SL bound to the whole position |
| `ParentOid` | uint64 | ties a TP/SL leg to its parent order |
| `Leg` | uint32 | which leg of a TP/SL bracket (legs share `Oid`+`Status`) |
| `Position` | PerpPosition | position as of the event |
| `Oid` | uint64 | order id |
| `Cloid` | string | client order id |
| `Hash` | string | L1 tx hash; empty if the order never hit chain |
| `PlacedTime` | int64 | order placement time, epoch **ms** |
| `Builder` | string | order builder address (usually empty) |
| `BuilderFee` | uint32 | builder fee rate in tenths of a basis point |
| `EventTime` | int64 | status-event time, epoch **ns** (vs `PlacedTime` = placement) |
| `Index` | uint32 | place of the status in the block's order-status list, from 1 — the event order; the legs of a bracket share it (`Leg` tells them apart); `0` on data written before the field existed |

An order can report two statuses in one block (`open`, then `filled` or `canceled`), and both
carry the **same** `EventTime`: only `Index` says which came first. The node lists them in event
order, so a later `Index` is the later status.

### 5.5 TWAP (`Twap`)

TWAP (time-weighted average price) parent-order lifecycle: large orders sliced into
sub-orders over `DurationMinutes`. The individual executions appear in `Trades` (with
`IsTwap=true` / `TwapId`, and an all-zero `Execution.Hash`); this list carries the parent
order's state. Attribute both sides of a TWAP fill via the trade-key triple, not via
`TwapId` (only the owner's side carries it).

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | TWAP owner |
| `Market` | Market | resolved market (perp / HIP-3 / spot) |
| `TwapId` | uint64 | TWAP order id |
| `Order` | TwapOrder | the request |
| `Interval` | TwapInterval | timing |
| `State` | TwapExecution | cumulative execution state |

**`TwapOrder`:** `Side` (`"Buy"`/`"Sell"`), `Size` (total requested), `ReduceOnly`,
`Randomize` (randomized sub-order timing).

**`TwapInterval`:** `StartTime` (TWAP start, epoch **ms**), `DurationMinutes` (uint32),
`EventTime` (status-event time, epoch **ns**).

**`TwapExecution`:** `Status`
(`"activated"` \| `"finished"` \| `"terminated"` \| `"stopped"` \| `"error"`),
`StatusError` (set when `Status="error"`), `ExecutedSize` (cumulative filled),
`ExecutedNotional` (cumulative, USD).

### 5.6 Book updates (`BookUpdate`)

Per-order (L4) order-book change — market microstructure, distinct from an order's
lifecycle in `OrderUpdate`. Sourced from the node `BookDiff` stream.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | owner of the resting order |
| `Market` | Market | resolved market |
| `Oid` | uint64 | order id |
| `Side` | string | `"Buy"` \| `"Sell"` |
| `Px` | string | price level |
| `Kind` | string | `"new"` \| `"update"` \| `"remove"` |
| `Size` | string | new size / size after |
| `SizeBefore` | string | size before (`update` only) |
| `Position` | PerpPosition | position as of the event |
| `Index` | uint32 | ordinal within the block |

### 5.7 Funding & leverage (`PerpFunding`, `TraderLeverageUpdate`)

**`PerpFunding`** — hourly funding, per (trader, market). Sourced from misc-events funding.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | position owner |
| `Market` | Market | perp market |
| `Amount` | string | signed (+ received, − paid) |
| `Rate` | string | signed hourly rate |
| `Size` | string | signed position size |

**`TraderLeverageUpdate`** — an accepted `updateLeverage`, from the **consensus** stream.
The leverage *change* itself; the resulting value is also on every `PerpPosition.Leverage`.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | trader |
| `Market` | Market | perp market |
| `Leverage` | uint32 | new leverage |
| `IsCross` | bool | `true` = cross, `false` = isolated |

### 5.8 Prices (`PriceUpdate`)

Oracle / mark price republishes, from the HIP-3 oracle stream flattened to one row per
`(Market, Kind)`.

| Field | Type | Description |
|---|---|---|
| `Market` | Market | resolved market |
| `Price` | string | the price |
| `Kind` | string | `"mark"` \| `"oracle"` \| `"markInput"` \| `"spotInput"` \| `"extPerpInput"` \| `"extPerp"` |
| `UpdateClass` | string | `"Normal"` \| `"Deployer"` \| `"Fallback"` |
| `DailyPx` | string | daily reference price |
| `UpdateTime` | int64 | epoch **ns** |

### 5.9 Signed consensus actions (`SignedActionEvent`)

One denormalized row per consensus action (from `replica_cmds`): block/bundle context +
the ecrecover'd `Signer` + parsed leverage. It carries what the event streams never do —
**rejects** (`Status="err"`) and the real `Signer`. **Only rejects and rare action types
ride this list**; a successful order/cancel already appears in `Orders`/`Trades` with more
structure. Addresses/hashes are `bytes` here (thin denormalization of the raw consensus row).

| Field | Type | Description |
|---|---|---|
| `Consensus` | SignedActionConsensus | round / proposer context |
| `Bundle` | SignedActionBundle | bundle context |
| `User` | bytes | acting **master** account (NOT the signer) |
| `Signer` | bytes | ecrecover'd agent wallet; empty for system/unsigned actions |
| `VaultAddress` | bytes | set when acting for a vault |
| `ActionType` | string | `"order"` \| `"cancel"` \| `"spotSend"` \| … (open set) |
| `Nonce` | uint64 | epoch **ms** by convention |
| `ExpiresAfter` | int64 | epoch-ms deadline, `0` when absent |
| `Status` | string | `"ok"` \| `"err"` |
| `Leverage` | SignedActionLeverage | set only for accepted `updateLeverage` |
| `Action` | bytes | raw JSON of the action, verbatim |
| `Response` | bytes | raw JSON: statuses/oids for `"ok"`, error text for `"err"` |

**`SignedActionConsensus`:** `Round`, `ParentRound`, `Proposer` (`0x…` validator that
proposed the block), `HardforkVersion`.

**`SignedActionBundle`:** `Hash` (L1 bundle hash — **disjoint** from `Fill.Hash` /
`OrderStatus.Hash`, §3), `Broadcaster`, `BroadcasterNonce`, `Index` (bundle ordinal within
block), `ActionIndex` (action ordinal within bundle).

**`SignedActionLeverage`:** `Asset` (asset id; `0` = BTC, valid), `Value` (leverage
multiplier), `IsCross`.

---

## 6. Candles — `CandlesBlockMessage`

`candles.proto`. Delivered alongside the block message: the OHLC bars touched by one trades
block, mirroring the envelope (`BlockNumber` + epoch-ns `BlockTime` + repeated payload).

| Message | Field | Type | Description |
|---|---|---|---|
| `CandlesBlockMessage` | `BlockNumber` | uint64 | canonical block id |
| | `BlockTime` | int64 | epoch **ns** |
| | `Candles` | repeated Candle | bars touched by this block |
| `Candle` | `Market` | Market | shared `Market` (§5.2) |
| | `Interval` | TimeInterval | bucket |
| | `Ohlc` | Ohlc | aggregates |
| `TimeInterval` | `Duration` | uint32 | bucket width, **seconds** — `60`/`300`/`900`/`3600` = 1m/5m/15m/1h |
| | `Start` | uint32 | bar open, epoch **seconds** (`toStartOfInterval`) |
| | `End` | uint32 | bar close boundary, epoch seconds (`Start + Duration`) |
| `Ohlc` | `Open`/`High`/`Low`/`Close`/`Volume` | double | Float64 — full precision even for high-priced assets |

---

## 7. State snapshots — `HyperliquidSnapshotMessage`

`hyperliquid_snapshot_message.proto`. The exchange **state** at one block, every 10,000 blocks
(≈ 12 min): what the event streams never carry — every open position with its margin, every
account's balance on every perp dex, the whole L4 book **with owners** plus the trigger orders
still waiting, spot balances, Earn (the borrow / lend pool), vaults and their depositors.

The source is the node's own state dump, `periodic_abci_states/<YYYYMMDD>/<height>.rmp`
(`hl-visor` writes it; msgpack, ~1.9 GB, root `exchange`). Bitquery archives every file byte
for byte, so **the `.rmp` is the raw layer** — there is no raw snapshot protobuf. Its keys are
mostly one letter and undocumented; the meanings below were checked against the node's info
API (`clearinghouseState`, `spotClearinghouseState`, `vaultDetails`), the HIP-3 oracle stream
and the node's own L4 dump (`hl-node compute-l4-snapshots --include-users
--include-trigger-orders`) (*observed*, October 2026).

All conventions of §2 hold: decimal strings, `0x…` strings for addresses, **empty = source
silent, NOT zero**. Rows are built like the block product (§5): the shared `Trader`, `Market`
and `PerpPosition` of §5.2, plus small per-topic sub-messages (as `Execution`, `Fees` or
`TwapOrder` there). `Market` is built exactly as for trades, so snapshot rows join trade rows
on `Market.CoinRaw`. `Trader` carries only `Address` — except on orders, whose `Signer`,
`Vault`, `Broadcaster` and `SignedAt` come from the processor's order attribution by `Oid`,
as on `OrderUpdate` (§7.5). A number or flag the snapshot may lack is `optional`: **null =
unknown, never 0** (strings keep empty = unknown). Fields marked *computed* are not in the
snapshot: they are derived from its other fields at the market's `MarkPx` **of the same
snapshot** (§7.2).

One converter makes every row, as a pure function of the `.rmp`: the realtime stream, the
archive datasets and any later historical export all come from the same code over the same
archived files, so they cannot disagree.

### 7.1 Envelope and parts

| Field | Type | Description |
|---|---|---|
| `BlockNumber` | uint64 | snapshot height (`exchange.locus.ctx.height`), a multiple of 10,000 |
| `BlockTime` | int64 | that block's consensus time, epoch **ns** (`ctx.time`) — equal to the block stream's `BlockTime` |
| `Part` | uint32 | `1..Parts` |
| `Parts` | uint32 | number of parts of this snapshot |
| `Markets` | repeated MarketState | every market: perp, HIP-3, spot, outcome (§7.3) |
| `Accounts` | repeated AccountState | per (account, perp dex) with collateral or a position (§7.4) |
| `Positions` | repeated PositionState | every open perp position (§7.4) |
| `LeverageSettings` | repeated LeverageSetting | per (account, perp market) the account touched, with a position or without (§7.4) |
| `Orders` | repeated OrderState | every order in the books and every waiting trigger order (§7.5) |
| `Balances` | repeated BalanceState | per (account, spot token), non-zero only (§7.6) |
| `EarnReserves` | repeated EarnReserveState | every token of Earn: pool totals, utilization, LTV, rate terms (§7.8) |
| `EarnPositions` | repeated EarnPositionState | per (account, token) with a supply or a borrow in Earn (§7.8) |
| `Vaults` | repeated VaultState | every vault (§7.7) |
| `VaultFollowers` | repeated VaultFollowerState | per (vault, depositor), leader included (§7.7) |

A snapshot is far larger than one message (hundreds of MB), so it is cut into `Parts`
messages of ≤ ~1 MB with the same `BlockNumber`/`BlockTime`. Rows go into parts in list order
(`Markets`, `Accounts`, `Positions`, `LeverageSettings`, `Orders`, `Balances`, `EarnReserves`,
`EarnPositions`, `Vaults`, `VaultFollowers`); a part may carry rows of several lists, and every
row is in exactly one part. A snapshot never changes, so a re-delivered part is identical — unlike block messages
(§9), **deduplicating on `(BlockNumber, Part)` is safe**. The snapshot is whole once all parts
`1..Parts` have arrived.

**Every list in every snapshot.** Each snapshot (each 10,000 blocks) carries all the lists, the heavy,
slow-moving ones too — leverage settings (~8.6 M rows), spot balances (~2.5 M), vault depositors
(~0.2 M) — so any snapshot answers any question about its height alone. Field 5 (`IsFull`) marked the
heights that are multiples of 300,000, once the only ones to carry those three lists; it is reserved.

Field numbers 14–17 of the envelope are kept for sections planned later: staking
(`c_staking`), sub-accounts (`sub_account_tracker`), agents (`agent_tracker`) and fee tracking
(`locus.ftr`); Earn, added after them, took 18 and 19.

### 7.2 Units and computed fields

How the snapshot stores numbers, and what the product turns them into:

| Quantity | Stored in the `.rmp` as | Product |
|---|---|---|
| perp size, open interest | integer = value × 10^`szDecimals` | decimal string |
| perp price (book, oracle, trigger) | integer = value × 10^(6 − `szDecimals`) | decimal string |
| spot / outcome price | integer = value × 10^(8 − `szDecimals` of the base token) | decimal string |
| perp-dex money: balance, entry notional, funding, isolated margin, tier bounds, vault deposits | integer micro-units (× 10^6) of the dex collateral, signed where it can be negative | decimal string |
| spot balance, hold | integer = value × 10^`weiDecimals` of the token | decimal string |
| spot entry notional | integer = USDC × 10^8 | decimal string |
| Earn shares and basis | integer = value × 10^`weiDecimals` of the token | decimal string |
| Earn indexes, rates, LTV, protocol share | integer × 10^14 | decimal string |
| time | ISO string with nanoseconds | epoch ns (order placement: the stored epoch-ms integer) |
| vault fraction, leader commission | float | shortest decimal string |
| perp asset index | native `i`; HIP-3 dex `n`: `n × 10000 + i` | `MarketState.AssetId` uses the action numbering: `100000 + n × 10000 + i` |

Decimal strings are printed the way the node prints them in its streams and in its L4 dump —
trailing zeros trimmed, a whole number keeps `.0` (`"82076.0"`, `"0.02986"`) — so a value
reads the same in `OrderState` and `OrderUpdate`.

The snapshot **omits zero numbers**: a number missing from an entry that exists (an account
with no cross balance, a balance with no entry notional) is zero and is written as `"0.0"`.
An empty field still means the source has no such value at all (`Prices.ExternalPerpPx` of a
spot market, `Tif` of a waiting trigger order).

The **mark price** is the second of the three prices the snapshot keeps per perp
(`cls[dex].oracle.pxs2[i]` = [oracle, mark, external perp], each `{u: time, p: price,
d: daily price}`; the order was checked against `PriceUpdate` of the same block).

Computed fields — `Mark` is the market's `Prices.MarkPx`; the sums run over one account's
positions on one dex:

| Field | Formula |
|---|---|
| `PositionState.Position.EntryPrice` | `Valuation.EntryNotional / abs(Size)` |
| `PositionState.Valuation.PositionValue` | `abs(Size) × Mark` |
| `PositionState.Valuation.UnrealizedPnl` | `Size × Mark − sign(Size) × EntryNotional` |
| `PositionState.Margin.MarginUsed` | cross: `PositionValue / Leverage`; isolated: `IsolatedRawUsd + Size × Mark` |
| `PositionState.Margin.MaintenanceMargin` | `PositionValue / (2 × MaxLeverage) − MaintenanceDeduction` of the last tier with `LowerBound ≤ PositionValue` |
| `AccountState.Cross.AccountValue` | `Cross.RawUsd + Σcross Size × Mark` |
| `AccountState.Total.AccountValue` | `Cross.AccountValue + Σisolated (IsolatedRawUsd + Size × Mark)` + isolated margin left in markets without a position |
| `AccountState.Total.RawUsd` | `Cross.RawUsd + Σisolated IsolatedRawUsd` + isolated margin left in markets without a position |
| `AccountState.{Total,Cross}.MarginUsed` | `Σ MarginUsed` — every position / cross only; `Total` also counts isolated margin left in markets without a position |
| `AccountState.{Total,Cross}.TotalNotional` | `Σ PositionValue` — every position / cross only |
| `AccountState.{Total,Cross}.MaintenanceMarginUsed` | `Σ MaintenanceMargin` — every position / cross only |
| `MarketState.Margin.IsolatedOnly` | `Mode` is `"strictIsolated"` or `"noCross"` |
| `MarketState.Margin.DefaultLeverage` | `min(20, Market.MaxLeverage)` |
| `LeverageSetting.EffectiveLeverage` | `Leverage`, else `min(20, MaxLeverage)`, else `min(20, Market.MaxLeverage)` |
| `LeverageSetting.EffectiveIsCross` | `IsCross`, else `not Margin.IsolatedOnly` |
| `VaultState.Summary.Equity` | the vault's `Total.AccountValue` on the native dex + its net USDC in Earn; a parent vault (HLP) + its children's equity (§7.7) |
| `VaultFollowerState.Share.Equity` / `AllTimePnl` | `Fraction × Summary.Equity` / `Equity − NetDeposit` |
| `AccountState` / `PositionState` / `OrderState` `.IsVault` | the account is one of `locus.vlt[]` |
| `MarketOutcome.SiblingCoin` | `#<encoding XOR 1>` |

These are the info API's `clearinghouseState` / `vaultDetails` fields (`positionValue`,
`unrealizedPnl`, `marginUsed`, `marginSummary`, `crossMarginSummary`,
`crossMaintenanceMarginUsed`, `vaultEquity`, `allTimePnl`). Sums are exact; the per-position
divisions — `MarginUsed` of a cross position, `MaintenanceMargin` — are **cut** toward zero to
6 decimals, and `EntryPrice` to the market's price decimals, as Hyperliquid does (rounding
missed `entryPx` on a quarter of the positions, and `crossMaintenanceMarginUsed` of a
50-position account by 0.000011).

*Observed* on snapshot 1194040000 (October 23, 2026) against the public info API, answers
taken just before and just after the block time and compared where the account did not
change in between: 190 (account, dex) pairs, 374 positions, 354 orders, 71 spot balances,
547 perp markets, 722 outcome sides, the 100 listed HLP depositors and 500 depositors of 5
more vaults matched field for field — sizes, entry prices, leverage, isolated margins,
cumulative funding, raw balances, every order field, and the computed fields to the last
digit when put at the API's own mark price.

The liquidation price and the withdrawable amount are not in the product yet. For the
withdrawable amount the documented rule ([margining](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/margining):
`transfer_margin_required = max(initial_margin_required, 0.1 × total_position_value)`) holds
with the dex's own share and the margin of open orders: `max(0, Cross.AccountValue −
max(Cross.MarginUsed, Cross.TotalNotional / max_withdraw_leverage) − order margin)`, where
`max_withdraw_leverage` is `cls[dex].max_withdraw_leverage` (10 native, 50 on the HIP-3 dexes)
and the order margin `cls[dex].user_states.open_order_tracker.user_to_data[account].e`. It
matched 383 of 402 answers; the rest were market makers whose orders moved in the minutes
between the snapshot and the answer, and unified / portfolio-margin accounts, whose spot side
the rule does not cover.

One market's oracle differed: `io:GPRO`, 1.018 here and in the node's own price stream
(`PriceUpdate` of block 1194040000: oracle 1.018, mark 1.0233), 1.0251 on the info API — the
API's `oraclePx` there is not the chain's oracle.

### 7.3 Markets (`MarketState`)

Source: perp — `exchange.locus.cls[dex].meta.universe[i]`, `.meta.marginTableIdToMarginTable`,
`.oracle.pxs2[i]`, `.user_states.asset_to_oi_szi`, `.halted_assets`; spot and outcome —
`exchange.locus.scl.meta.spot_infos[i]` / `token_infos`, prediction markets also
`scl.meta.outr`; the book from `exchange.perp_dexs[dex].books[i]` / `exchange.spot_books[i]`.
`Margin`, `Prices` and the open interest are perp-only and stay empty on spot and outcome
markets; `Outcome` is set on outcome markets only.

| Field | Type | Description |
|---|---|---|
| `Market` | Market | resolved market (§5.2) |
| `AssetId` | uint32 | asset id of exchange actions ([asset IDs](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api/asset-ids)): perp `i`, HIP-3 `100000 + dex × 10000 + i`, spot `10000 + i`, outcome `100000000 + encoding`, where `encoding = 10 × outcome + side` is the `n` of `CoinRaw` `#<n>` |
| `SzDecimals` | uint32 | size decimals; spot/outcome: the base token's |
| `Margin` | MarketMargin | margin mode, collateral, tiers |
| `Prices` | MarketPrices | oracle, mark and external perp prices |
| `Book` | MarketBook | top of the book |
| `OpenInterest` | string | one side, in coins (info API `openInterest`) |
| `Positions` | uint32 | *computed*: open positions in the market |
| `IsHalted` | bool | trading halted |
| `Outcome` | MarketOutcome | prediction (outcome) markets only: the question and the side |

**`MarketMargin`:**

| Field | Type | Description |
|---|---|---|
| `Mode` | string | `"normal"` \| `"strictIsolated"` \| `"noCross"` |
| `Collateral` | string | the dex's collateral token: `"USDC"` \| `"USDH"` \| `"USDE"` \| `"USDT0"` \| … |
| `Tiers` | repeated MarginTier | by `LowerBound`; one tier when the market has no tier table |
| `IsolatedOnly` | bool | *computed*: no cross margin in this market — an account that never set leverage here is isolated |
| `DefaultLeverage` | uint32 | *computed*: the leverage of an account that never set one here, `min(20, Market.MaxLeverage)` |

**`MarginTier`** — from `LowerBound` of position notional on, `MaxLeverage` applies and the
maintenance margin is `notional / (2 × MaxLeverage) − MaintenanceDeduction`
([margin tiers](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/margin-tiers)). A
market whose `marginTableId` is below 50 has no table: one tier, `MaxLeverage` = the id.

| Field | Type | Description |
|---|---|---|
| `LowerBound` | string | USD |
| `MaxLeverage` | uint32 | max leverage in the tier |
| `MaintenanceDeduction` | string | USD |

**`MarketPrices`:**

| Field | Type | Description |
|---|---|---|
| `OraclePx` | string | oracle price |
| `MarkPx` | string | mark price — the price of every *computed* field |
| `ExternalPerpPx` | string | external perp price |
| `MarkDailyPx` | string | daily reference price of the mark (as `PriceUpdate.DailyPx`) |
| `UpdateTime` | int64 | last oracle update, epoch **ns** |

**`MarketBook`:**

| Field | Type | Description |
|---|---|---|
| `LastTradePx` | string | last trade price |
| `BestBid` / `BestAsk` | string | best price levels; empty = that side is empty |
| `Orders` | uint32 | *computed*: orders resting in the book |

**`MarketOutcome`** — one side of a prediction market, from `scl.meta.outr` (`ots[]` per
outcome, `qts[]` per question). It is field for field the public info API `outcomeMeta`
(*observed*: the same 341 outcomes and 32 questions, the same names and descriptions). An
outcome has two sides, each its own market (§4): `encoding = 10 × OutcomeId + Side`, `CoinRaw`
`#<encoding>`, token `+<encoding>`. In the snapshot a side is a spot pair whose base token is
the side's token (`ots[].s[].t`), so its book is that pair's `spot_books[pair]`; the pair is
listed once, as the outcome market, not also as an `@<pair>` spot market. The price of a side
in its `Book` reads as the probability, 0..1.

| Field | Type | Description |
|---|---|---|
| `OutcomeId` | uint32 | outcome id (`outcomes[].outcome`) |
| `Side` | uint32 | `0` \| `1` |
| `SideName` | string | `sideSpecs[Side].name`: `"Yes"` \| `"No"` \| … |
| `SiblingCoin` | string | the other side of the same outcome, `#<encoding XOR 1>` |
| `Name` | string | outcome name, e.g. `"template:sportsTournamentParticipant"` |
| `Description` | string | `key:value\|key:value` as published, e.g. `"participant:Oklahoma City Thunder"` |
| `Question` | OutcomeQuestion | the question the outcome answers; empty when it stands alone |

**`OutcomeQuestion`** — a question groups several outcomes, one per answer
(`namedOutcomes`), plus a fallback ("other"):

| Field | Type | Description |
|---|---|---|
| `QuestionId` | uint32 | question id (`questions[].question`) |
| `Name` | string | e.g. `"template:sportsContestResult"` |
| `Description` | string | `key:value\|…` as published: competition, participants, resolution deadline, … |
| `IsFallback` | bool | this outcome is the question's `fallbackOutcome` |
| `IsSettled` | bool | this outcome is in `settledNamedOutcomes` |

### 7.4 Accounts and positions (`AccountState`, `PositionState`)

Source: `exchange.locus.cls[dex].user_states.user_to_state[]` — `[address, state]`, one
clearinghouse per perp dex (the native one and every HIP-3 dex keep separate margin, each in its
own collateral). The state's `u` is the cross balance; `p.p[]` holds one `[asset, entry]` per
asset the account touched — only entries with a size are positions, the others keep a leverage
setting.

**`AccountState`** — one account on one dex (info API `clearinghouseState` with that `dex`),
for every account with a cross balance, a position, or isolated margin left in a market:

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | the account (`Address` only) |
| `Protocol` | string | `"hypercore"` native, else the HIP-3 dex — as `Market.Protocol` |
| `Collateral` | string | the dex's collateral token |
| `Total` | MarginSummary | every position (`marginSummary`) |
| `Cross` | MarginSummary | cross positions only (`crossMarginSummary`) |
| `Positions` | uint32 | *computed*: open positions on this dex |
| `IsVault` | bool | *computed*: the account is a vault (`locus.vlt[]`) — HLP or a user vault |

**`MarginSummary`** — the info API's `marginSummary` / `crossMarginSummary`, plus the
maintenance margin, which the API gives for cross only (`crossMaintenanceMarginUsed`):

| Field | Type | Description |
|---|---|---|
| `AccountValue` | string | *computed* (`accountValue`) |
| `MarginUsed` | string | *computed* (`totalMarginUsed`) |
| `MaintenanceMarginUsed` | string | *computed*: `Σ MaintenanceMargin` (`Cross`: `crossMaintenanceMarginUsed`) |
| `TotalNotional` | string | *computed* (`totalNtlPos`) |
| `RawUsd` | string | `Cross`: the cross balance `u`, signed — collateral net of the cross positions' entry notional; `Total`: *computed*, plus the isolated margins (`totalRawUsd`) |

**`PositionState`** — one open position (info API `assetPositions[].position`):

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | position owner (`Address` only) |
| `Market` | Market | perp market |
| `Position` | PerpPosition | the shared position object of trades (§5.2) |
| `Margin` | PositionMargin | isolated margin, margin used, maintenance margin |
| `Valuation` | PositionValuation | entry notional and the position marked |
| `CumFunding` | CumFunding | cumulative funding, as on the Hyperliquid position |
| `IsVault` | bool | *computed*: the owner is a vault — HLP or a user vault |

**`PerpPosition` in a snapshot** — the same object as on trades, filled from the snapshot:

| Field | Description |
|---|---|
| `Side` | `"Long"` \| `"Short"` |
| `Size` | signed (`szi`) |
| `EntryPrice` | *computed* (`entryPx`) |
| `Leverage` | the setting; an entry without one is at the exchange default `min(20, MaxLeverage)` |
| `IsCross` | cross vs isolated; an entry without a setting is cross unless the market is isolated-only |
| `SizeBefore`, `RealizedPnl`, `Funding`, `Closed` | empty — event fields; funding is in `PositionState.CumFunding` |

**`PositionMargin`:**

| Field | Type | Description |
|---|---|---|
| `IsolatedRawUsd` | string | isolated only, signed: margin net of the entry notional (`leverage.rawUsd`) |
| `MarginUsed` | string | *computed* (`marginUsed`) |
| `MaintenanceMargin` | string | *computed* |

**`PositionValuation`:**

| Field | Type | Description |
|---|---|---|
| `EntryNotional` | string | USD paid at entry |
| `MarkPx` | string | the market's `Prices.MarkPx` in this snapshot |
| `PositionValue` | string | *computed* (`positionValue`) |
| `UnrealizedPnl` | string | *computed* (`unrealizedPnl`) |

**`CumFunding`** — cumulative funding of the position as Hyperliquid keeps it: the info API
`clearinghouseState` `cumFunding`, the same values, signed **+ paid by the account,
− received**. That is the **opposite** sign to `PerpFunding.Amount` (§5.7, + received): the
funding events of a position since it opened sum to `−SinceOpen` (*observed*: a short received
funding at 11:00 on October 22 — `PerpFunding.Amount` +0.097984 — and its `cumFunding.allTime`
fell by the same 0.097984).

| Field | Type | Description |
|---|---|---|
| `AllTime` | string | over the account's whole history in the market |
| `SinceOpen` | string | since the position opened |
| `SinceChange` | string | since its size last changed |

**`LeverageSetting`** — every `[asset, entry]` of `p.p[]`, with a
position or without (~8.6 M on mainnet, October 2026; open positions are ~0.4 M of them): the
leverage the account's **next** trade in that market gets, until an accepted `updateLeverage`
(`TraderLeverageUpdate`, §5.7) changes it. A leverage set without a position counts too. A
market the account never touched has no row and is at `MarketMargin.DefaultLeverage` /
`IsolatedOnly`. With the stream's leverage changes, the earliest snapshot is the starting point
for the leverage of any account at any later block.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | the account (`Address` only) |
| `Market` | Market | perp market |
| `Leverage` | optional uint32 | the account's own setting (`l.C` cross / `l.I.l` isolated); null = never set |
| `IsCross` | optional bool | null = never set |
| `MaxLeverage` | optional uint32 | the entry's own cap: its margin table (`M`) resolved to the first tier's max leverage — it can differ from the market's (HLP vaults hold every market at 50); null = no cap of its own |
| `EffectiveLeverage` | uint32 | *computed*: the setting, else the default under the entry's cap, else under the market's (§7.2) |
| `EffectiveIsCross` | bool | *computed*: the setting, else cross unless the market is isolated-only |

The same rule fills `PositionState.Position.Leverage` / `IsCross`, and matched the info API on
sampled accounts, HLP vaults included (*observed*).

### 7.5 Orders (`OrderState`)

Source: `exchange.perp_dexs[dex].books[i]` and `exchange.spot_books[i]`, one book per market.
The book proper is a slab `bod.e[]` of orders linked per price level (`halfs[side].levels[]`
names the first slot of each level, every order the next); orders are listed in that queue
order. Trigger orders that have not fired are not in the book: they wait in the book's
`user_states[].U[]`. The rows match the node's own L4 dump, owner included (*observed*: the
same order counts on a full mainnet snapshot, sampled orders field by field). The row is flat,
like `OrderUpdate` (§5.4), with the same names and values.

The snapshot names only the owner. Who signed the order is not in it: `Trader.Signer`,
`Vault`, `Broadcaster` and `SignedAt` come from the processor's order attribution — the same
map of live orders by `Oid` that fills them on `OrderUpdate` — and stay empty for an order it
does not know.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | owner (`Address`) and its signer (`Signer`, `Vault`, `Broadcaster`, `SignedAt`) |
| `Market` | Market | resolved market |
| `Side` | string | `"Buy"` \| `"Sell"` |
| `OrderType` | string | `"Limit"` \| `"Stop Market"` \| `"Stop Limit"` \| `"Take Profit Market"` \| `"Take Profit Limit"` \| `"Trailing Stop Market"` |
| `Tif` | string | `"Gtc"` \| `"Alo"`; empty on waiting trigger orders and on the oldest book orders |
| `ReduceOnly` | bool | order may only reduce a position |
| `LimitPx` | string | limit price (for a market trigger: its slippage limit) |
| `Size` | string | **remaining** size; `"0.0"` on a position TP/SL, which closes the whole position |
| `OrigSz` | string | original size |
| `TriggerPx` | string | trigger price, `"0.0"` without a trigger; trailing stop: *computed* from `Trailing.BestPx` and the retracement |
| `TriggerCondition` | string | `"N/A"` \| `"Price above 86000"` \| `"Triggered"` (fired, now in the book) \| trailing: `"Activation above 0.05, retracement 5%, best waiting"` |
| `IsTrigger` | bool | a trigger order still waiting — **not in the book** |
| `IsPositionTpsl` | bool | TP/SL bound to the whole position |
| `ParentOid` | optional uint64 | a TP/SL leg → its parent order, while the parent rests unfilled; null = none |
| `SiblingOid` | optional uint64 | the other leg of a TP/SL pair (one cancels the other); null = none |
| `Oid` | uint64 | order id |
| `Cloid` | string | client order id |
| `PlacedTime` | int64 | placement time, epoch **ms** |
| `Builder` | string | builder address (usually empty) |
| `BuilderFee` | uint32 | builder fee rate in tenths of a basis point |
| `QueuePosition` | optional uint32 | place at its price level, from `1` = filled first; null on waiting trigger orders |
| `Trailing` | TrailingStop | trailing stops only |
| `IsVault` | bool | *computed*: the owner is a vault — e.g. the HLP market-making orders |

**`TrailingStop`:**

| Field | Type | Description |
|---|---|---|
| `RetracementPct` | string | percent (`"15"` = 15%); empty when given as `RetracementPx` |
| `RetracementPx` | string | price distance; empty when given as `RetracementPct` |
| `ActivationPx` | string | empty = active from placement |
| `BestPx` | string | best price since activation; empty = not active yet |

### 7.6 Spot balances (`BalanceState`)

Source: `exchange.locus.scl.user_states[]` — `[address, {b: [[token, {t, h, e}]]}]` (info API
`spotClearinghouseState`). Zero balances are left out; token names and decimals come from
`scl.meta.token_infos`.

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | the account (`Address` only) |
| `Token` | string | `"USDC"`, `"HYPE"`, …; an outcome side's token is `"+<encoding>"`, as the info API names it |
| `TokenIndex` | uint32 | token index (spotMeta `tokens[].index`) |
| `Total` | string | balance — the spot ledger itself |
| `Hold` | string | part of `Total` locked in open orders |
| `EntryNotional` | string | USDC paid for `Total` |

`Total` is what the spot ledger holds. For an account that trades perps from its spot balance
(unified account, dex abstraction), the info API's `spotClearinghouseState` adds the USDC its
perp dexes hold to `total` and shows it as `hold`; that money is the account's
`AccountState.Total.AccountValue` here, so it is not counted twice (*observed*: the ledger
equals the API's `total` on plain accounts and `total − hold` on unified ones). The system
accounts that mint outcome tokens can go negative: the snapshot keeps that as a wrapped
unsigned integer and the API prints it as such (`184467440736536.5`); here it is the signed
value (`-559.00001`).

There is no USD value of a balance: Hyperliquid publishes none, and a price rule of our own
would not be its data. USD is priced downstream, by the Bitquery price index used for the other
products.

### 7.7 Vaults (`VaultState`, `VaultFollowerState`)

Source: `exchange.locus.vlt[]` — `[vault, {name, leader, user_states, …}]`. A vault trades as
its own account: its positions, margin and orders are the rows of that address in the lists
above. Depositors are their own list, since the largest vault has ~94,000 of them.

**`VaultState`:**

| Field | Type | Description |
|---|---|---|
| `Vault` | string | vault account |
| `Name` | string | |
| `Description` | string | |
| `Leader` | string | leader address |
| `CreateTime` | int64 | epoch **ns** |
| `Terms` | VaultTerms | commission, lockup, deposits |
| `Relationship` | VaultRelationship | parent / child vaults |
| `Summary` | VaultSummary | equity and depositors |

**`VaultTerms`:**

| Field | Type | Description |
|---|---|---|
| `LeaderCommission` | string | share of the followers' profit, e.g. `"0.1"` |
| `LockupSeconds` | uint32 | lockup after each deposit, **s** |
| `AllowDeposits` | bool | |
| `AlwaysCloseOnWithdraw` | bool | |
| `IsClosed` | bool | |

**`VaultRelationship`:** `Type` (`"normal"` \| `"parent"` \| `"child"`), `ChildVaults`
(repeated string, parent vaults only).

**`VaultSummary`:**

| Field | Type | Description |
|---|---|---|
| `Equity` | string | *computed*: the vault's `Total.AccountValue` on the native dex + its net USDC in Earn; a parent vault adds its children's |
| `LeaderFraction` | string | *computed*: the leader's share of the vault |
| `Followers` | uint32 | *computed*: depositors, leader included |

A vault's `Equity` is its account value on the native dex **plus what it holds in Earn**
(`locus.blp`, the borrow/lend pool: shares × the token's interest index / 10¹⁴, supply minus
borrow), and a **parent** vault adds its children's equity. That is how Hyperliquid values
HLP, which "supplies USDC in Earn" ([protocol vaults](https://hyperliquid.gitbook.io/hyperliquid-docs/hypercore/vaults/protocol-vaults)):
at 1194040000 HLP Strategy X held 100,598,280.50 USDC there, and HLP came to 177,129,170.86 —
the info API said 177,129,122.25 27 s later. A depositor's share is `Fraction × Equity`
(*observed*: the 100 HLP depositors the API lists within 4·10⁻⁷, and 500 depositors of the 5
largest open vaults — shares, lockups and the leader's fraction exact). A vault with Earn in a
token other than USDC would need a price: its `Equity` (and its depositors' `Share.Equity` /
`AllTimePnl`) stays empty — none at 1194040000.

**`VaultFollowerState`** — one depositor (info API `vaultDetails.followers`):

| Field | Type | Description |
|---|---|---|
| `Vault` | string | vault account |
| `Trader` | Trader | the depositor (`Address` only) |
| `IsLeader` | bool | the depositor is the leader |
| `Share` | VaultShare | share, equity, deposits, PnL |
| `Times` | VaultFollowTimes | entry, last deposit, end of lockup |

**`VaultShare`:**

| Field | Type | Description |
|---|---|---|
| `Fraction` | string | share of the vault |
| `Equity` | string | *computed* (`vaultEquity`) |
| `NetDeposit` | string | deposits − withdrawals, USD, signed |
| `AllTimePnl` | string | *computed* (`allTimePnl`) |

**`VaultFollowTimes`:**

| Field | Type | Description |
|---|---|---|
| `EntryTime` | int64 | first deposit, epoch **ns** (`vaultEntryTime`) |
| `LastDepositTime` | int64 | last deposit, epoch **ns** |
| `LockupUntil` | int64 | *computed*: `LastDepositTime + Terms.LockupSeconds`, epoch **ns** (`lockupUntil`) |

### 7.8 Earn (`EarnReserveState`, `EarnPositionState`)

Earn is Hyperliquid's borrow / lend pool: accounts supply tokens and earn interest, and
portfolio-margin accounts borrow from it ([portfolio margin](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/portfolio-margin));
HLP supplies USDC there (§7.7). Source: `exchange.locus.blp` — `r.t[]` one entry per token
(USDC, HYPE, UBTC, USDT0, USDH at 1194040000), `u[]` one entry per account. A position is held
as shares of the pool; its value with interest is `shares × index / 10^14`, the token's index
rising hourly. Both lists ride every snapshot (5 tokens, ~17,000 positions, ~2 MB).

**`EarnReserveState`** — one token (info API `allBorrowLendReserveStates`):

| Field | Type | Description |
|---|---|---|
| `Token` / `TokenIndex` | string / uint32 | the token |
| `TotalSupplied` | string | *computed*: the pool's supply shares (`l[1].t`) × `SupplyIndex` (`totalSupplied`) |
| `TotalBorrowed` | string | *computed*: the pool's borrow shares (`l[0].t`) × `BorrowIndex` (`totalBorrowed`) |
| `Utilization` | string | *computed*: `TotalBorrowed / TotalSupplied`, cut to 10 decimals (`utilization`) |
| `Ltv` | string | share of the token's value that can be borrowed against it (`v`, `ltv`): HYPE `0.65`, UBTC `0.5`; `"0.0"` = not collateral |
| `SupplyIndex` / `BorrowIndex` | string | interest indexes (`l[1].i` / `l[0].i`): what one share is worth in tokens |
| `UpdateTime` | int64 | the last interest update (`T`), epoch **ns** |
| `RateModel` | EarnRateModel | the terms the yearly rates follow from |

**`EarnRateModel`** (`y`, and `r.p` for the protocol share):

| Field | Type | Description |
|---|---|---|
| `BaseRate` | string | yearly borrow rate up to the kink (`y.i`), `0.05` |
| `KinkUtilization` | string | the kink (`y.u`), `0.8` |
| `MaxRate` | string | yearly borrow rate at full utilization (`y.a`), `1.0` |
| `ProtocolShare` | string | share of the borrow interest the protocol keeps as a liquidation buffer (`r.p`), `0.1` |

The yearly rates are not in the row — they follow from it, so ClickHouse / GraphQL compute
them (as an `ALIAS`), and a change of Hyperliquid's formula needs no reprocessing:

```
BorrowYearlyRate = BaseRate + (MaxRate − BaseRate) / (1 − KinkUtilization) × max(0, Utilization − KinkUtilization)
SupplyYearlyRate = BorrowYearlyRate × Utilization × (1 − ProtocolShare)
```

For stablecoins that is the documented `0.05 + 4.75 × max(0, utilization − 0.8)`. With the API's
own utilization it gave the API's `borrowYearlyRate` and `supplyYearlyRate` to the last of their
10 digits for all five tokens (*observed*). `y.k` (`0.05`, equal to `y.i` on every token) is
not exposed: its role below the kink cannot be told apart while it equals the base rate.

**`EarnPositionState`** — one account's position in one token (info API `borrowLendUserState`
`tokenToState`):

| Field | Type | Description |
|---|---|---|
| `Trader` | Trader | the account (`Address` only) |
| `Token` / `TokenIndex` | string / uint32 | the token |
| `Supply` / `Borrow` | EarnAmount | the two sides, `"0.0"` where the account has none |
| `IsVault` | bool | *computed*: the account is a vault — HLP Strategy X supplies ~100.6 M USDC |

**`EarnAmount`:**

| Field | Type | Description |
|---|---|---|
| `Basis` | string | what was supplied / borrowed, without interest (`b`, `basis`) |
| `Value` | string | *computed*: `Shares × index / 10^14`, with interest (`value`) |
| `Shares` | string | shares of the pool (`s`) |

*Observed* against the info API 1.8 h after snapshot 1194040000: `Basis` equal on 146 of 154
sampled positions (the other 8 had supplied or borrowed since), `Value` higher by the interest
of those 1.8 h only (1.0·10⁻⁵ on a 5% borrow), LTVs equal, pool totals within the flows since.

---

## 8. Raw source streams (node ingestion — internal)

The raw layer that the product (§5–§6) is projected from. A **lossless typed re-encoding**
of the node NDJSON — documented here for reference; subscribers consume the product, not
these.

Summary of all eight raw streams:

| # | Proto | Message | Node source (flag → path under `~/hl/data`) |
|---|---|---|---|
| 1 | `fills.proto` | `Fill` | `--write-fills` → `node_fills_streaming/hourly/…` |
| 2 | `order_statuses.proto` | `OrderStatus` | `--write-order-statuses` → `node_order_statuses_streaming/hourly/…` |
| 3 | `raw_book_diffs.proto` | `BookDiff` | `--write-raw-book-diffs` → `node_raw_book_diffs_streaming/hourly/…` |
| 4 | `twap_statuses.proto` | `TwapStatus` | (with fills) → `node_twap_statuses_streaming/hourly/…` |
| 5 | `oracle_updates.proto` | `OracleUpdate` | `--write-hip3-oracle-updates` → `hip3_oracle_updates_streaming/hourly/…` |
| 6 | `misc_events.proto` | `MiscEvent` | `--write-misc-events` → `misc_events_streaming/hourly/…` |
| 7 | `core_writer_actions.proto` | `SystemAction` | → `system_and_core_writer_actions_streaming/hourly/…` |
| 8 | `transactions.proto` | `HyperCoreTransactionsBlock` | `replica_cmds/{session}/{date}/{height}` |

Streams 1–7 are aggregated per block into `HyperCoreBlock`; stream 8 has its own container.

Every raw event stream line carries the same envelope, which maps to the block header:

```json
{"local_time":"…","block_time":"2026-07-17T11:39:05.530222877","block_number":1075904092,"events":[ … ]}
```

`local_time` (the node's own wall clock) is intentionally **not** modeled — it is
node-specific, not consensus data.

Depending on the node's write flags (`--batch-by-block` vs `--stream-with-block-info`),
one block of one stream is either a single line or **many consecutive lines** — each
repeating the same envelope and carrying a slice of the block's events (*observed*: a
single mainnet block spread over 3,782 order-status lines). The protobuf
`HyperCoreBlock` always aggregates the whole block regardless of the node's write mode.

---

### 8.1 Fills (`fills.proto` — `Fill`)

One **execution for one user**. A match normally yields two `Fill`s — taker
(`Crossed=true`) and maker (`Crossed=false`) — paired by `(BlockNumber, Coin, Tid)`.
Projected to `Trade` / `PerpLiquidation` (§5.3).

Real node line (mainnet):

```json
["0xf1e79e5cf063073deec9a0d86c178f7cde84d479",
 {"coin":"xyz:NFLX","px":"66.362","sz":"1.911","side":"B","time":1784288345530,
  "startPosition":"751.363","dir":"Open Long","closedPnl":"0.0",
  "hash":"0xce5d…0411","oid":497705647404,"crossed":false,"fee":"0.003804",
  "tid":870392364981122,"cloid":"0x0000…9087","feeToken":"USDC",
  "twapId":null,"deployerFee":"0.001902"}]
```

| Field | Type | Description |
|---|---|---|
| `User` | bytes | address of **this side** of the match |
| `Coin` | string | market id (see §4) |
| `Side` | string | `"B"` = bid (this user bought), `"A"` = ask (sold) |
| `Px` | string | execution price |
| `Sz` | string | executed size (base units of the market) |
| `StartPosition` | string | user's signed position in this market **before** the fill |
| `Dir` | string | position-change label, open set: `"Open Long"`, `"Close Short"`, `"Buy"`, `"Sell"`, `"Long > Short"`, `"Split Outcome"`, `"Merge Outcome"`, … |
| `ClosedPnl` | string | realized PnL of the position part closed by this fill |
| `Fee` | string | total fee paid by this side, denominated in `FeeToken`; negative = rebate |
| `FeeToken` | string | fee denomination — an **open set**, not always `"USDC"`: spot fills may pay in the spot token (`HYPE`, `PURR`, `UBTC`, …), outcome markets in `+<n>` (~3.7% of fills are non-USDC, *observed*) |
| `Hash` | bytes | L1 transaction hash — synthetic, embeds (block, action seq, sub-order), see §3. Both sides of a match carry the **taker's** transaction hash. **Not unique** (one tx → many fills) and **all-zero (32 zero bytes) on both sides of every TWAP fill** (~11.6% of all fills, *observed*). Never a dedup/join key |
| `Oid` | uint64 | order id of this side's order |
| `Tid` | uint64 | 50-bit match id shared by both sides — see §3 |
| `Crossed` | bool | `true` = this side was the taker/aggressor |
| `Time` | int64 | event time, epoch **ms** |
| `Cloid` | optional bytes | client order id; the node sends both `null` and key-absent — both map to unset |
| `TwapId` | optional uint64 | set **only on the TWAP owner's side** of a TWAP fill; the counterparty has the same all-zero `Hash` but no `TwapId` (*observed*: exactly 74,530 vs 74,530 in a 1.28M-fill sample). Attribute both sides to a TWAP via the trade-key triple, not via this field |
| `BuilderFee` | optional string | portion of `Fee` forwarded to the builder (same unit as `Fee`). Never present without `Builder`; the reverse does happen (~16% of builder fills carry no `BuilderFee`, *observed*). Not a fixed share: `BuilderFee/Fee` *observed* between 0.45 and 1.00 |
| `Liquidation` | optional FillLiquidation | present only when this fill was forced by a liquidation |
| `Builder` | optional bytes | builder address ([builder codes](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/builder-codes) — per-order fee attribution to the app that routed it) |
| `DeployerFee` | optional string | portion of `Fee` to the HIP-3 asset deployer, same unit as `Fee` (*observed* share of `Fee`: 0.5 dominant, then 0.45) |
| `PriorityGas` | optional string | priority-gas payment. **Not** a component of `Fee` and **not** denominated in `FeeToken` — one fixed unit for all fills (which one, USDC or HYPE, is not derivable from the data). May be literally `"0.0"` |
| `Extra` | map | forward-compat net (§2.4) |

`FillLiquidation`:

| Field | Type | Description |
|---|---|---|
| `LiquidatedUser` | bytes | address being liquidated (may be this fill's own `User` — closing their position — or the counterparty) |
| `MarkPx` | string | mark price at liquidation |
| `Method` | string | `"market"` \| `"backstop"` (only `"market"` *observed* so far) |

**Fee algebra:** `BuilderFee` and `DeployerFee` are *components* of `Fee` (same unit,
already included). `PriorityGas` is *separate* — do not add it to `Fee`.

---

### 8.2 Order statuses (`order_statuses.proto` — `OrderStatus`)

One order-lifecycle event: placement, fill, cancel, trigger, or any of the many
rejections. This stream is the **complete L1 order audit trail** — including orders that
never reached the book. Projected to `OrderUpdate` (§5.4).

| Field | Type | Description |
|---|---|---|
| `User` | bytes | order owner |
| `Status` | string | open set: `"open"`, `"filled"`, `"canceled"`, `"triggered"`, `"rejected"`, and many specific variants *observed* like `"badAloPxRejected"`, `"iocCancelRejected"`, `"reduceOnlyRejected"`, `"perpMarginRejected"`, `"marginCanceled"`, … |
| `Time` | int64 | event time, epoch **ns** |
| `Order` | Order | full order snapshot at event time |
| `Hash` | optional bytes | L1 tx hash; `null` for order events with no direct transaction |
| `Builder` | optional OrderBuilder | builder attribution, usually `null` |
| `Extra` | map | §2.4 |

`Order` (also used recursively for `Children`):

| Field | Type | Description |
|---|---|---|
| `Coin` | string | market id (§4) |
| `Side` | string | `"B"` bid / `"A"` ask |
| `LimitPx` | string | limit price |
| `Sz` | string | **remaining** size |
| `OrigSz` | string | original size |
| `Oid` | uint64 | order id |
| `Timestamp` | int64 | order placement time, epoch **ms** |
| `OrderType` | string | open set: `"Limit"`, `"Market"`, `"Stop Market"`, `"Stop Limit"`, `"Take Profit Market"`, `"Take Profit Limit"` |
| `TriggerCondition` | string | `"N/A"` or a human-readable condition, e.g. `"Price above 917.08"` |
| `TriggerPx` | string | trigger price (`"0.0"` when not a trigger order) |
| `IsTrigger` | bool | true for stop / take-profit orders |
| `IsPositionTpsl` | bool | true when the TP/SL is bound to the whole position rather than a specific order |
| `ReduceOnly` | bool | order may only reduce a position |
| `Cloid` | optional bytes | client order id, may be `null` |
| `Tif` | optional string | time-in-force: `"Gtc"`, `"Alo"` (add-liquidity-only / post-only), `"Ioc"`, `"FrontendMarket"`, `"LiquidationMarket"`; `null` on trigger orders |
| `Children` | repeated Order | TP/SL bracket attached to a parent order — see below |
| `Extra` | map | §2.4 |

**Children — TP/SL brackets** (*audited on 3M orders*): a bracket holds 1 child (SL only)
or 2 (TP + SL). A child's field set is exactly `Order` (verified: zero extra, zero missing
keys); nesting never goes deeper than one level. Children are always `ReduceOnly`, have
`Tif=null`, and `OrderType` ∈ {`Stop Market`, `Take Profit Market`, `Stop Limit`,
`Take Profit Limit`}.

`OrderBuilder` — node shape `{"b": <address>, "f": <fee>}`:

| Field | Type | Description |
|---|---|---|
| `Address` | bytes | builder address (node key `"b"`) |
| `Fee` | uint32 | builder fee rate in tenths of a basis point (node key `"f"`) |

---

### 8.3 Raw book diffs (`raw_book_diffs.proto` — `BookDiff`)

One **L4 order-book delta** — order-level (not price-level-aggregated) book changes.
Applied on top of an L4 snapshot, this stream reconstructs the entire book in real time,
with every resting order attributed to its owner. Projected to `BookUpdate` (§5.6).

L4 snapshots (the anchor state) come from the node itself: offline via
`hl-node --chain <chain> compute-l4-snapshots <abci-state-path> <out-path>` (state files
are written every 10,000 blocks to `~/hl/data/periodic_abci_states/{date}/{height}.rmp`),
or online from a node running with `--serve-info` via a
`{"type":"fileSnapshot","request":{"type":"l4Snapshots",…}}` request.

The node's `raw_book_diff` field is a tagged union, flattened into `Kind` + variant fields:

| Node JSON | `Kind` | Variant fields set |
|---|---|---|
| `{"new":{"sz":…}}` | `"new"` | `Sz` — size resting at the level |
| `{"update":{"origSz":…,"newSz":…}}` | `"update"` | `OrigSz`, `NewSz` — size before / after |
| `"remove"` | `"remove"` | none |

| Field | Type | Description |
|---|---|---|
| `User` | bytes | owner of the resting order |
| `Oid` | uint64 | order id |
| `Coin` | string | market id (§4) |
| `Side` | string | `"B"` bid / `"A"` ask |
| `Px` | string | price level |
| `Kind` | string | union tag verbatim — an unseen variant still comes through |
| `Sz`, `OrigSz`, `NewSz` | optional string | per-variant, see table above |
| `Extra` | map | §2.4 |

---

### 8.4 TWAP statuses (`twap_statuses.proto` — `TwapStatus`)

Lifecycle of TWAP (time-weighted average price) orders: large orders sliced into
sub-orders over `Minutes`. The individual executions appear in the **fills** stream
(all-zero `Hash`, `TwapId` on the owner side); this stream carries the parent order's
state transitions. Projected to `Twap` (§5.5).

The node's `status` is `string | {"error": msg}`, flattened:

| Field | Type | Description |
|---|---|---|
| `Time` | int64 | event time, epoch **ns** |
| `TwapId` | uint64 | TWAP order id (matches `Fill.TwapId`) |
| `State` | TwapState | full state snapshot at the event |
| `Status` | string | `"activated"` \| `"finished"` \| `"terminated"` \| `"error"` |
| `StatusError` | optional string | set when status was `{"error": …}`, e.g. `"Insufficient margin to place order."` |
| `Extra` | map | §2.4 |

`TwapState`:

| Field | Type | Description |
|---|---|---|
| `Coin` | string | market — perps, HIP-3 perps and spot (`@<n>`) all appear |
| `User` | bytes | TWAP owner |
| `Side` | string | `"B"` / `"A"` |
| `Sz` | string | total order size |
| `ExecutedSz` | string | filled so far |
| `ExecutedNtl` | string | executed notional (USD) |
| `Minutes` | uint32 | TWAP duration |
| `Timestamp` | int64 | order **start** time, epoch **ms** — *not* the event time (on `"finished"` events it is minutes earlier than `Time`) |
| `ReduceOnly` | bool | |
| `Randomize` | bool | randomized sub-order timing |
| `Extra` | map | §2.4 |

---

### 8.5 HIP-3 oracle updates (`oracle_updates.proto` — `OracleUpdate`)

Price-oracle updates for [HIP-3](https://hyperliquid.gitbook.io/hyperliquid-docs/hips/hip-3-builder-deployed-perpetuals)
builder-deployed perp dexes (equities, FX, commodities, indices — `xyz:AAPL`,
`mkts:GOLD`, `flx:BTC`, …). HIP-3 deployers push their own oracle prices; this stream
records each push and the resulting oracle state. Projected to `PriceUpdate` (§5.8).

| Field | Type | Description |
|---|---|---|
| `UpdateClass` | string | `"Deployer"` (pushed by the dex deployer) \| `"Fallback"` |
| `MarkPxInputs` | repeated PxInput | submitted mark-price inputs, `[coin, px]` pairs |
| `SpotPxInputs` | repeated PxInput | submitted spot-price inputs |
| `ExternalPerpPxInputs` | repeated PxInput | external perp prices; often empty |
| `OraclePxs` | OraclePxs | resulting oracle state |
| `Extra` | map | §2.4 |

`PxInput` = `{Coin, Px}` — one `[coin, price]` pair.

`OraclePxs` — three coin→price lists, each entry an `OraclePx`:

| Field | Type | Description |
|---|---|---|
| `CoinToMarkPx` | repeated OraclePx | mark prices |
| `CoinToOraclePx` | repeated OraclePx | oracle prices |
| `CoinToExternalPerpPx` | repeated OraclePx | external perp prices |

`OraclePx`:

| Field | Type | Description |
|---|---|---|
| `Coin` | string | market id |
| `Px` | string | current price |
| `LastUpdateTime` | int64 | epoch **ns** |
| `DailyPx` | string | daily reference price |

---

### 8.6 Misc events (`misc_events.proto` — `MiscEvent`)

Everything financial that is not a trade: ledger movements (transfers, deposits,
withdrawals, vault operations), hourly funding, staking, delegation, and a few consensus
events. The most polymorphic stream. Funding feeds the product's `PerpFunding` list (§5.7);
the rest is carried on the raw layer.

Envelope: `{time, hash, inner:{"<Variant>":{…}}}` → `InnerType` = the variant tag;
exactly one variant field is populated.

| Field | Type | Description |
|---|---|---|
| `Time` | int64 | event time, epoch **ns** |
| `Hash` | bytes | L1 tx hash; **all-zero for system-generated events** (funding, validator rewards, …) |
| `InnerType` | string | `"LedgerUpdate"` \| `"Funding"` \| `"CDeposit"` \| `"CWithdrawal"` \| `"Delegation"` \| `"ValidatorRewards"` \| `"GossipPriorityAuctionRestart"` |
| one of the variant fields below | | |
| `Extra` | map | §2.4 |

#### LedgerUpdate — balance changes

`{Users: [addresses touched], Delta: LedgerDelta}`. The node's `delta` is a tagged union
of ~16 shapes flattened into one message: `Type` is the tag, and each type populates only
its own subset of fields (everything else unset).

| Field | Type | Used by (examples) |
|---|---|---|
| `Type` | string | `"send"`, `"spotTransfer"`, `"deposit"`, `"withdraw"`, `"borrowLend"`, `"accountClassTransfer"`, `"cStakingTransfer"`, `"vaultWithdraw"`, … (open set) |
| `User` | optional bytes | initiator |
| `Amount`, `Token` | optional string | amount and token of the movement |
| `Usdc`, `UsdcValue` | optional string | USDC amount / USD value |
| `Destination` | optional bytes | recipient (`send`, transfers) |
| `DestinationDex`, `SourceDex` | optional string | dex namespaces for cross-dex transfers (empty string = default dex, `"spot"` = spot balance) |
| `Fee`, `FeeToken`, `NativeTokenFee` | optional string | transfer fees |
| `Nonce` | optional uint64 | user action nonce |
| `ToPerp` | optional bool | `accountClassTransfer`: direction spot↔perp |
| `Operation`, `InterestAmount` | optional string | `borrowLend` |
| `Vault`, `Basis`, `ClosingCost`, `Commission`, `NetWithdrawnUsd`, `RequestedUsd` | optional | vault operations (`vaultWithdraw`, …) |
| `IsDeposit` | optional bool | `cStakingTransfer`: direction |

Real example (`type:"send"`):

```json
{"LedgerUpdate":{"users":["0xf70d…dbef","0xe609…ab29"],
 "delta":{"type":"send","user":"0xf70d…dbef","destination":"0xe609…ab29",
          "sourceDex":"","destinationDex":"spot","token":"USDC",
          "amount":"33.053999","usdcValue":"33.053999","fee":"0.0",
          "nativeTokenFee":"0.0","nonce":1784288317818,"feeToken":""}}}
```

#### Funding — hourly funding payments

`Funding.Deltas` — one `FundingDelta` per (user, coin) with an open position:

| Field | Type | Description |
|---|---|---|
| `User` | bytes | position owner |
| `Coin` | string | perp market |
| `FundingAmount` | string | signed payment (positive = received) |
| `Szi` | string | signed position size at funding |
| `FundingRate` | string | applied hourly rate |

Funding events are block-wide and huge — a single funding line can exceed 20 MB of JSON
(*observed*).

#### Staking & consensus variants

* `CDeposit` `{Amount, User}` — HYPE moved into staking ("C" = core staking balance).
* `CWithdrawal` `{Amount, User, IsFinalized}` — staking withdrawal (unstaking has an
  unbonding delay; `IsFinalized` marks completion).
* `Delegation` `{Amount, User, Validator, IsUndelegate}` — (un)delegation to a validator.
* `ValidatorRewards` — per-validator reward list; payload kept raw in `Extra`
  (`validator_to_reward`: list of `[validator, amount]`).
* `GossipPriorityAuctionRestart` `{SlotId}` + `Extra` (`end_gas`, `previous_winner`) —
  consensus-level gossip priority auction event.

---

### 8.7 Core-writer & system actions (`core_writer_actions.proto` — `SystemAction`)

The **HyperEVM → HyperCore bridge**: actions on the Core side that originate from EVM
transactions — system transfers of tokens/USD between the two sides, and Core actions
(orders, agent approvals, delegations) submitted by EVM contracts through the CoreWriter
system contract. Every event links to its EVM transaction via `EvmTxHash`.

> ⚠ Unlike every other stream, `Token` / `Wei` / `Ntl` here are **raw integers**
> (token index, wei amount, notional), not decimal strings — this is the node's own
> format for this stream.

| Field | Type | Description |
|---|---|---|
| `User` | bytes | acting address (system addresses like `0x2000…00c5` appear for system transfers) |
| `Nonce` | uint64 | event nonce |
| `EvmTxHash` | bytes | hash of the HyperEVM transaction that caused this Core action |
| `Action` | Action | flattened action union |
| `Extra` | map | §2.4 |

`Action` — `Type` is the tag (open set); each type populates its own fields:

| `Type` | Fields set | Meaning |
|---|---|---|
| `"SystemSpotSendAction"` | `Destination`, `Token`, `Wei` | spot token transfer EVM→Core |
| `"SystemSendAssetAction"` | `Destination`, `Token`, `Wei`, `DestinationDexOrSpot`, `SourceDexOrSpot`, `FromSubAccount`* | asset transfer between dex/spot balances |
| `"SystemUsdClassTransferAction"` | `Ntl`, `ToPerp` | USD transfer spot↔perp class |
| `"order"` | `Grouping`, `Orders`, opt. `Builder`, `MaxFeeRate` | order placement via CoreWriter. Each `Orders` entry is the raw JSON of one order in the exchange wire format — keys `a` (asset index), `b` (isBuy), `p` (price), `s` (size), `r` (reduceOnly), `t` (type, e.g. `{"limit":{"tif":"Ioc"}}`), `c` (cloid) |
| `"approveAgent"` | `AgentAddress`, `AgentName`, `SignatureChainId`, `HyperliquidChain`, `ActionNonce` | approve an agent (API wallet) — signed L1 action bridged from EVM |
| `"tokenDelegate"` | `Validator`, `Wei`, `IsUndelegate`, `SignatureChainId`, `HyperliquidChain`, `ActionNonce` | (un)delegate stake to a validator |

\* `FromSubAccount`: always null in samples so far — type tentative.

`ActionNonce` is the action-level nonce of the bridged signed action and is **not** the
same as the event-level `Nonce`.

---

### 8.8 Consensus transactions (`transactions.proto` — replica_cmds)

The input side of the chain: every signed action bundle that consensus ordered into a
block, with the execution engine's verdict. Source: the node's `replica_cmds` files
(one line = one whole block; a node restart starts a new `{session}` directory which may
replay heights). The product surfaces the useful subset (rejects, real signer, leverage)
through `SignedActionEvent` (§5.9); this raw container carries the full block.

Design choice: the **skeleton is typed** (block header, signatures, nonces, verdict) while
the action payload — an open, hardfork-evolving set of dozens of action types — is carried
**verbatim as raw JSON bytes** (`Action`, `Response`). A typed copy would silently drop
unknown fields; parse at read time instead.

`HyperCoreTransactionsBlock`:

| Field | Type | Description |
|---|---|---|
| `BlockNumber` | uint64 | canonical block id (not present in the raw line — derived from file name + line index) |
| `BlockTime` | int64 | consensus time, epoch **ns** — same clock as `HyperCoreBlock` |
| `Round`, `ParentRound` | uint64 | consensus round counters; **≠ BlockNumber** (empty rounds produce no block; `Round` runs ~28% ahead, *observed*) |
| `Proposer` | bytes | validator that proposed the block |
| `HardforkVersion` | uint32 | hardfork version at this height |
| `Bundles` | repeated TransactionBundle | the block's transaction bundles |

`TransactionBundle`:

| Field | Type | Description |
|---|---|---|
| `Hash` | bytes | bundle hash as recorded in replica_cmds. **Not** equal to `Fill.Hash` / `OrderStatus.Hash` (verified disjoint) — join the views via the oid route or `BlockNumber` + `User` (§3) |
| `Broadcaster` | bytes | node/relay that broadcast the bundle |
| `BroadcasterNonce` | uint64 | |
| `Actions` | repeated SignedAction | |

`SignedAction`:

| Field | Type | Description |
|---|---|---|
| `User` | bytes | the acting **master** account as resolved by the engine. **Not the signer**: an agent (API wallet) signs on the master's behalf; the actual signer is recoverable via ecrecover from `SignatureR/S/V` + `Action` |
| `ActionType` | string | `action.type` verbatim: `"order"`, `"cancel"`, `"spotSend"`, `"tokenDelegate"`, `"evmRawTx"`, … (open set) |
| `SignatureR`, `SignatureS` | bytes | 32-byte signature halves |
| `SignatureV` | uint32 | recovery id, 27/28 |
| `Nonce` | uint64 | signer nonce (epoch ms by wallet convention) |
| `VaultAddress` | bytes | set when acting on behalf of a vault/subaccount, else empty |
| `ExpiresAfter` | int64 | epoch-ms deadline, 0 = absent |
| `Status` | string | engine verdict: `"ok"` \| `"err"`. **Rejections that never reach the event streams (bad nonce, API errors) surface only here** |
| `Response` | bytes | raw JSON engine response: statuses/oids for `"ok"`, the error text for `"err"` |
| `Action` | bytes | raw JSON of the action, verbatim |

---

## 9. Delivery semantics (Bitquery pipeline)

For consumers of the Kafka / archive representation of these messages:

1. A message covers **one (stream, block)** and is always a **full snapshot** from the
   start of the block — never an increment.
2. Delivery is **at-least-once**; a redelivery is a superset of what was seen; the last
   version wins. **Deduplicating by block number alone is therefore forbidden** — a later
   message for the same block may contain more events.
3. Assembling all streams of one block is the consumer's job; there is no
   "block complete" signal.
4. Archive objects pack exactly 1000 consecutive blocks; blocks that are empty in the
   event streams still appear (with the real `BlockTime`) so every bucket holds 1000
   entries.

Snapshots (§7) are the exception to 1–2: a snapshot is the state at one block, cut into
`Parts` messages that never change, so a redelivered part is identical and
`(BlockNumber, Part)` is a safe dedup key; the snapshot is complete once parts `1..Parts`
have arrived.

---

## 10. Generated code

* Go: `messages/*.pb.go` (package `hyperliquid_messages`), one file per proto — the product
  in `hyperliquid_block_message.pb.go`, `candles.pb.go` and `hyperliquid_snapshot_message.pb.go`
* Python: `python/hyperliquid/*_pb2.py`, likewise (`hyperliquid_snapshot_message_pb2.py`)
* The protos carry no `option go_package`: every proto needs its own
  `--go_opt="Mhyperliquid/<file>.proto=hyperliquid/messages;hyperliquid_messages"` line in
  the root `Makefile`.
* Regenerate: `make generate_hyperliquid` from the repository root.

## 11. Official references

* [Node repository & data flags](https://github.com/hyperliquid-dex/node)
* [L1 data schemas](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/nodes/l1-data-schemas)
* [Historical data (S3 buckets)](https://hyperliquid.gitbook.io/hyperliquid-docs/historical-data)
* [Exchange API (order wire format, statuses)](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api)
* [HIP-3: builder-deployed perpetuals](https://hyperliquid.gitbook.io/hyperliquid-docs/hips/hip-3-builder-deployed-perpetuals)
* [Builder codes](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/builder-codes)
* [Margining](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/margining),
  [margin tiers](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/margin-tiers),
  [liquidations](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/liquidations) — the formulas of §7.2
* [Info endpoint, perpetuals](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api/info-endpoint/perpetuals) — `clearinghouseState`, `metaAndAssetCtxs`
* [Asset IDs](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api/asset-ids) — perp, HIP-3, spot and outcome numbering (`MarketState.AssetId`)
