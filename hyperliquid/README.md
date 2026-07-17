# Hyperliquid (HyperCore) Protobuf Schemas

Canonical protobuf schemas for Bitquery's Hyperliquid **HyperCore** streaming pipeline.

HyperCore is the native execution engine of the Hyperliquid L1 — the on-chain perpetual
and spot exchange (order books, matching, liquidations, funding, staking, vaults). It is
distinct from **HyperEVM**, the EVM side of the same chain; the two are linked, and one of
the streams below (`CoreWriterActions`) records exactly that link.

Every message here mirrors data written by the official Hyperliquid node
(`hl-visor run-non-validator`) as NDJSON files under `~/hl/data`. The schemas are a
**lossless, typed re-encoding** of those files: nothing is aggregated, filtered or
re-interpreted. Where the node's JSON uses tagged unions, the schema flattens them into a
tag field plus optional variant fields; where a field is unknown or unstable, it is kept
verbatim in an `Extra` map or as raw JSON bytes, so no information is ever dropped.

**Schema evolution.** When the node starts emitting a new field, it first appears in
`Extra` (§2.4) — already usable, as a raw JSON string. The upgrade path is: add a typed
`optional` field under a **fresh** field number, and from that release on the value
arrives typed. Historical messages are **not rewritten** — they keep the value in
`Extra` under the node's original field name. A consumer that needs full history reads
the typed field and falls back to `Extra["<name>"]` for older blocks; the two encodings
never overlap for the same block, and nothing is lost in either period. Proto field
numbers are never reused or renumbered, so old and new binaries stay compatible both ways.

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
 ├── 7 event streams   <stream>_streaming/hourly/{YYYYMMDD}/{hour}   1 line = 1 block, or a slice of one
 └── replica_cmds/{session}/{date}/{height-base}                     1 line = 1 whole consensus block
        │
        ▼  typed protobuf encoding (this package)
 ├── HyperCoreBlock  / HyperCoreBlocks              (hypercore.proto)     — the 7 event streams
 └── HyperCoreTransactionsBlock / …Blocks           (transactions.proto)  — consensus transactions
```

Two complementary views of the chain:

| View | Container | Contents | Question it answers |
|---|---|---|---|
| **Events** | `HyperCoreBlock` | What the execution engine **produced** in a block: fills, order status changes, book deltas, TWAP lifecycle, oracle prices, ledger/funding/staking events, EVM↔Core bridge actions | "What happened?" |
| **Transactions** | `HyperCoreTransactionsBlock` | What went **into** the block: every signed action bundle consensus ordered, with signatures, nonces and the engine's accept/reject verdict | "Who asked for it, and was it accepted?" |

Rejections that never produce events (bad nonce, malformed action, insufficient margin at
the API level) are visible **only** in the transactions view.

### Blocks

* HyperCore has **no block hash**. `BlockNumber` is the canonical block identifier.
* `BlockNumber ↔ BlockTime` is a bijection — no two blocks share a timestamp
  (*observed*: 374,720 distinct pairs in 374,720 blocks over 12 h).
* `BlockTime` is the consensus timestamp in **epoch nanoseconds (UTC)**, identical across
  all streams of the same block.
* Blocks are fast: roughly 10–15 blocks/s on mainnet (*observed*, July 2026 — block
  numbers were around 1.07–1.08 billion).
* A block in which a given stream had no events produces **no line** in that stream's file.
  Blocks empty across all 7 event streams are ~5% of all heights (*observed*).
* `Round` / `ParentRound` (transactions view only) are consensus counters, **not** block
  numbers: empty rounds produce no block, so `Round` runs ahead of `BlockNumber`
  (~28% ahead, *observed*).

### File / message layout

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

The one deliberate exception: `core_writer_actions.proto`, where `Token`, `Wei` and `Ntl`
are raw integers (a token index, a wei amount, an integer notional), because that is what
the node itself emits for this stream.

### 2.2 Timestamps are mixed — per field, not per stream

The node mixes epoch-millisecond integers and ISO nanosecond strings. The schemas keep
each field's native precision (ISO strings are parsed to integer nanoseconds):

| Field | Unit | Meaning |
|---|---|---|
| `HyperCoreBlock.BlockTime`, `HyperCoreTransactionsBlock.BlockTime` | **ns** | consensus block time |
| `Fill.Time` | **ms** | fill event time |
| `OrderStatus.Time` | **ns** | status event time |
| `Order.Timestamp` | **ms** | order placement time |
| `TwapStatus.Time` | **ns** | TWAP event time |
| `TwapState.Timestamp` | **ms** | TWAP order **start** time (not event time) |
| `MiscEvent.Time` | **ns** | event time |
| `OraclePx.LastUpdateTime` | **ns** | last oracle update |
| `SignedAction.Nonce` | ms *(by wallet convention)* | signer nonce |
| `SignedAction.ExpiresAfter` | ms | action deadline, `0` = absent |

### 2.3 Addresses and hashes are `bytes`

`0x…` hex strings from the node are decoded to raw bytes: 20 bytes for addresses,
32 bytes for hashes. Empty/absent → empty bytes or unset `optional`.

### 2.4 `Extra` maps — the forward-compatibility net

Nearly every message carries `map<string,string> Extra`. Any JSON key the node emits that
the schema does not model lands there — key = the node's field name, value = the raw JSON
value. In the common case `Extra` is empty; a non-empty `Extra` means the node started
emitting a new field and nothing was lost.

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
* One match normally produces **two** fills — one per counterparty.

| Identifier | What it is | Unique? | Safe usage |
|---|---|---|---|
| `BlockNumber` | canonical block id | globally unique | primary block key |
| `Oid` | exchange-assigned order id (uint64) | unique per order | order lifecycle joins |
| `Cloid` | optional client-supplied order id (16 bytes) | client-controlled | client-side matching only |
| `Tid` | **50-bit** hash of `(buyer_oid, seller_oid)` | **no** — collides at scale (~2²⁵ trades) | unique only within one block+coin |
| `Hash` | L1 transaction hash | **no** — shared by many fills; **all-zero on both sides of TWAP fills** | display / grouping, never dedup |
| `TwapId` | TWAP order id | unique per TWAP | see caveat in §4.1 |
| **`(BlockNumber, Coin, Tid)`** | **the trade key** | **yes** (*observed* on millions of fills) | dedup, pairing the two sides of a match |

**Event hashes are synthetic.** `Fill.Hash` / `OrderStatus.Hash` / `MiscEvent.Hash` is
*not* a hash of the transaction content. Starting at byte 10 it embeds a compact
length-prefixed big-endian encoding of the event's position:

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

**Joining events to transactions:** the bundle `Hash` recorded in `replica_cmds` lives in
a different universe than event hashes (*verified disjoint*). The precise join is
three-step (*verified on a full mainnet block*):

1. `Fill.Hash` equals the `Hash` of the aggressor order's `open`/`filled` `OrderStatus`
   in the same block — both sides of a match carry the **taker's** transaction hash
   (34/34 fills in the verified block).
2. That `OrderStatus` carries the `Oid`.
3. The `Oid` appears in the engine response (`SignedAction.Response`, from the node's
   `resps`) of exactly one signed action (1,424/1,437 `open` statuses in the verified
   block) — giving the precise bundle and action.

TWAP fills (all-zero hash) attribute via the trade-key triple instead (§5.1). Where the
oid route is unavailable (e.g. rejected actions that never produced an order id), fall
back to the coarse join `BlockNumber` + `User`.

---

## 4. Market (coin) symbology

The `Coin` field is a raw market identifier with four families:

| Pattern | Family | Example | Notes |
|---|---|---|---|
| plain symbol | native perpetual | `BTC`, `HYPE`, `kBONK` | `k`-prefix = 1000× denomination (e.g. `kBONK` = 1000 BONK) |
| `<dex>:<SYM>` | HIP-3 (builder-deployed) perpetual | `xyz:SP500`, `flx:BTC`, `mkts:AAPL` | prefix = the deployer's dex namespace; equities/FX/commodities live here |
| `@<n>` | spot pair by index | `@107` | n = spot pair index in exchange metadata |
| `#<n>` | outcome token (prediction market) | `#1890` | comes in complementary pairs `(n, n^1)` — see below |

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

## 5. Stream reference

Summary of all eight streams:

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

### 5.1 Fills (`fills.proto` — `Fill`)

One **execution for one user**. A match normally yields two `Fill`s — taker
(`Crossed=true`) and maker (`Crossed=false`) — paired by `(BlockNumber, Coin, Tid)`.

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

### 5.2 Order statuses (`order_statuses.proto` — `OrderStatus`)

One order-lifecycle event: placement, fill, cancel, trigger, or any of the many
rejections. This stream is the **complete L1 order audit trail** — including orders that
never reached the book.

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

### 5.3 Raw book diffs (`raw_book_diffs.proto` — `BookDiff`)

One **L4 order-book delta** — order-level (not price-level-aggregated) book changes.
Applied on top of an L4 snapshot, this stream reconstructs the entire book in real time,
with every resting order attributed to its owner.

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

### 5.4 TWAP statuses (`twap_statuses.proto` — `TwapStatus`)

Lifecycle of TWAP (time-weighted average price) orders: large orders sliced into
sub-orders over `Minutes`. The individual executions appear in the **fills** stream
(all-zero `Hash`, `TwapId` on the owner side); this stream carries the parent order's
state transitions.

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

### 5.5 HIP-3 oracle updates (`oracle_updates.proto` — `OracleUpdate`)

Price-oracle updates for [HIP-3](https://hyperliquid.gitbook.io/hyperliquid-docs/hips/hip-3-builder-deployed-perpetuals)
builder-deployed perp dexes (equities, FX, commodities, indices — `xyz:AAPL`,
`mkts:GOLD`, `flx:BTC`, …). HIP-3 deployers push their own oracle prices; this stream
records each push and the resulting oracle state.

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

### 5.6 Misc events (`misc_events.proto` — `MiscEvent`)

Everything financial that is not a trade: ledger movements (transfers, deposits,
withdrawals, vault operations), hourly funding, staking, delegation, and a few consensus
events. The most polymorphic stream.

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

### 5.7 Core-writer & system actions (`core_writer_actions.proto` — `SystemAction`)

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

### 5.8 Consensus transactions (`transactions.proto` — replica_cmds)

The input side of the chain: every signed action bundle that consensus ordered into a
block, with the execution engine's verdict. Source: the node's `replica_cmds` files
(one line = one whole block; a node restart starts a new `{session}` directory which may
replay heights).

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

## 6. Delivery semantics (Bitquery pipeline)

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

---

## 7. Generated code

* Go: `messages/*.pb.go` (package `hyperliquid_messages`)
* Python: `python/`
* Regenerate: `make generate_hyperliquid` from the repository root.

## 8. Official references

* [Node repository & data flags](https://github.com/hyperliquid-dex/node)
* [L1 data schemas](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/nodes/l1-data-schemas)
* [Historical data (S3 buckets)](https://hyperliquid.gitbook.io/hyperliquid-docs/historical-data)
* [Exchange API (order wire format, statuses)](https://hyperliquid.gitbook.io/hyperliquid-docs/for-developers/api)
* [HIP-3: builder-deployed perpetuals](https://hyperliquid.gitbook.io/hyperliquid-docs/hips/hip-3-builder-deployed-perpetuals)
* [Builder codes](https://hyperliquid.gitbook.io/hyperliquid-docs/trading/builder-codes)
