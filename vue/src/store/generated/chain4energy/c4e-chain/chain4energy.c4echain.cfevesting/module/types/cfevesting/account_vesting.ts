/* eslint-disable */
import * as Long from "long";
import { util, configure, Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "chain4energy.c4echain.cfevesting";

export interface AccountVestingsList {
  vestings: AccountVestings[];
}

export interface AccountVestings {
  address: string;
  delegableAddress: string;
  /** uint64 delegated = 12; */
  vestings: Vesting[];
}

export interface Vesting {
  id: number;
  vestingType: string;
  vestingStartBlock: number;
  lockEndBlock: number;
  vestingEndBlock: number;
  vested: string;
  /**
   * uint64 claimable = 6;
   * int64 last_freeing_block = 7;
   */
  freeCoinsBlockPeriod: number;
  /** uint64 free_coins_per_period = 9; */
  delegationAllowed: boolean;
  withdrawn: string;
}

const baseAccountVestingsList: object = {};

export const AccountVestingsList = {
  encode(
    message: AccountVestingsList,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.vestings) {
      AccountVestings.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestingsList {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestingsList } as AccountVestingsList;
    message.vestings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.vestings.push(
            AccountVestings.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountVestingsList {
    const message = { ...baseAccountVestingsList } as AccountVestingsList;
    message.vestings = [];
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(AccountVestings.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AccountVestingsList): unknown {
    const obj: any = {};
    if (message.vestings) {
      obj.vestings = message.vestings.map((e) =>
        e ? AccountVestings.toJSON(e) : undefined
      );
    } else {
      obj.vestings = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AccountVestingsList>): AccountVestingsList {
    const message = { ...baseAccountVestingsList } as AccountVestingsList;
    message.vestings = [];
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(AccountVestings.fromPartial(e));
      }
    }
    return message;
  },
};

const baseAccountVestings: object = { address: "", delegableAddress: "" };

export const AccountVestings = {
  encode(message: AccountVestings, writer: Writer = Writer.create()): Writer {
    if (message.address !== "") {
      writer.uint32(10).string(message.address);
    }
    if (message.delegableAddress !== "") {
      writer.uint32(18).string(message.delegableAddress);
    }
    for (const v of message.vestings) {
      Vesting.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): AccountVestings {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseAccountVestings } as AccountVestings;
    message.vestings = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.address = reader.string();
          break;
        case 2:
          message.delegableAddress = reader.string();
          break;
        case 3:
          message.vestings.push(Vesting.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
    message.vestings = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (
      object.delegableAddress !== undefined &&
      object.delegableAddress !== null
    ) {
      message.delegableAddress = String(object.delegableAddress);
    } else {
      message.delegableAddress = "";
    }
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(Vesting.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: AccountVestings): unknown {
    const obj: any = {};
    message.address !== undefined && (obj.address = message.address);
    message.delegableAddress !== undefined &&
      (obj.delegableAddress = message.delegableAddress);
    if (message.vestings) {
      obj.vestings = message.vestings.map((e) =>
        e ? Vesting.toJSON(e) : undefined
      );
    } else {
      obj.vestings = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<AccountVestings>): AccountVestings {
    const message = { ...baseAccountVestings } as AccountVestings;
    message.vestings = [];
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (
      object.delegableAddress !== undefined &&
      object.delegableAddress !== null
    ) {
      message.delegableAddress = object.delegableAddress;
    } else {
      message.delegableAddress = "";
    }
    if (object.vestings !== undefined && object.vestings !== null) {
      for (const e of object.vestings) {
        message.vestings.push(Vesting.fromPartial(e));
      }
    }
    return message;
  },
};

const baseVesting: object = {
  id: 0,
  vestingType: "",
  vestingStartBlock: 0,
  lockEndBlock: 0,
  vestingEndBlock: 0,
  vested: "",
  freeCoinsBlockPeriod: 0,
  delegationAllowed: false,
  withdrawn: "",
};

export const Vesting = {
  encode(message: Vesting, writer: Writer = Writer.create()): Writer {
    if (message.id !== 0) {
      writer.uint32(8).int32(message.id);
    }
    if (message.vestingType !== "") {
      writer.uint32(18).string(message.vestingType);
    }
    if (message.vestingStartBlock !== 0) {
      writer.uint32(24).int64(message.vestingStartBlock);
    }
    if (message.lockEndBlock !== 0) {
      writer.uint32(32).int64(message.lockEndBlock);
    }
    if (message.vestingEndBlock !== 0) {
      writer.uint32(40).int64(message.vestingEndBlock);
    }
    if (message.vested !== "") {
      writer.uint32(50).string(message.vested);
    }
    if (message.freeCoinsBlockPeriod !== 0) {
      writer.uint32(56).int64(message.freeCoinsBlockPeriod);
    }
    if (message.delegationAllowed === true) {
      writer.uint32(64).bool(message.delegationAllowed);
    }
    if (message.withdrawn !== "") {
      writer.uint32(74).string(message.withdrawn);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Vesting {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseVesting } as Vesting;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.id = reader.int32();
          break;
        case 2:
          message.vestingType = reader.string();
          break;
        case 3:
          message.vestingStartBlock = longToNumber(reader.int64() as Long);
          break;
        case 4:
          message.lockEndBlock = longToNumber(reader.int64() as Long);
          break;
        case 5:
          message.vestingEndBlock = longToNumber(reader.int64() as Long);
          break;
        case 6:
          message.vested = reader.string();
          break;
        case 7:
          message.freeCoinsBlockPeriod = longToNumber(reader.int64() as Long);
          break;
        case 8:
          message.delegationAllowed = reader.bool();
          break;
        case 9:
          message.withdrawn = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Vesting {
    const message = { ...baseVesting } as Vesting;
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = String(object.vestingType);
    } else {
      message.vestingType = "";
    }
    if (
      object.vestingStartBlock !== undefined &&
      object.vestingStartBlock !== null
    ) {
      message.vestingStartBlock = Number(object.vestingStartBlock);
    } else {
      message.vestingStartBlock = 0;
    }
    if (object.lockEndBlock !== undefined && object.lockEndBlock !== null) {
      message.lockEndBlock = Number(object.lockEndBlock);
    } else {
      message.lockEndBlock = 0;
    }
    if (
      object.vestingEndBlock !== undefined &&
      object.vestingEndBlock !== null
    ) {
      message.vestingEndBlock = Number(object.vestingEndBlock);
    } else {
      message.vestingEndBlock = 0;
    }
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = String(object.vested);
    } else {
      message.vested = "";
    }
    if (
      object.freeCoinsBlockPeriod !== undefined &&
      object.freeCoinsBlockPeriod !== null
    ) {
      message.freeCoinsBlockPeriod = Number(object.freeCoinsBlockPeriod);
    } else {
      message.freeCoinsBlockPeriod = 0;
    }
    if (
      object.delegationAllowed !== undefined &&
      object.delegationAllowed !== null
    ) {
      message.delegationAllowed = Boolean(object.delegationAllowed);
    } else {
      message.delegationAllowed = false;
    }
    if (object.withdrawn !== undefined && object.withdrawn !== null) {
      message.withdrawn = String(object.withdrawn);
    } else {
      message.withdrawn = "";
    }
    return message;
  },

  toJSON(message: Vesting): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.vestingType !== undefined &&
      (obj.vestingType = message.vestingType);
    message.vestingStartBlock !== undefined &&
      (obj.vestingStartBlock = message.vestingStartBlock);
    message.lockEndBlock !== undefined &&
      (obj.lockEndBlock = message.lockEndBlock);
    message.vestingEndBlock !== undefined &&
      (obj.vestingEndBlock = message.vestingEndBlock);
    message.vested !== undefined && (obj.vested = message.vested);
    message.freeCoinsBlockPeriod !== undefined &&
      (obj.freeCoinsBlockPeriod = message.freeCoinsBlockPeriod);
    message.delegationAllowed !== undefined &&
      (obj.delegationAllowed = message.delegationAllowed);
    message.withdrawn !== undefined && (obj.withdrawn = message.withdrawn);
    return obj;
  },

  fromPartial(object: DeepPartial<Vesting>): Vesting {
    const message = { ...baseVesting } as Vesting;
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    if (object.vestingType !== undefined && object.vestingType !== null) {
      message.vestingType = object.vestingType;
    } else {
      message.vestingType = "";
    }
    if (
      object.vestingStartBlock !== undefined &&
      object.vestingStartBlock !== null
    ) {
      message.vestingStartBlock = object.vestingStartBlock;
    } else {
      message.vestingStartBlock = 0;
    }
    if (object.lockEndBlock !== undefined && object.lockEndBlock !== null) {
      message.lockEndBlock = object.lockEndBlock;
    } else {
      message.lockEndBlock = 0;
    }
    if (
      object.vestingEndBlock !== undefined &&
      object.vestingEndBlock !== null
    ) {
      message.vestingEndBlock = object.vestingEndBlock;
    } else {
      message.vestingEndBlock = 0;
    }
    if (object.vested !== undefined && object.vested !== null) {
      message.vested = object.vested;
    } else {
      message.vested = "";
    }
    if (
      object.freeCoinsBlockPeriod !== undefined &&
      object.freeCoinsBlockPeriod !== null
    ) {
      message.freeCoinsBlockPeriod = object.freeCoinsBlockPeriod;
    } else {
      message.freeCoinsBlockPeriod = 0;
    }
    if (
      object.delegationAllowed !== undefined &&
      object.delegationAllowed !== null
    ) {
      message.delegationAllowed = object.delegationAllowed;
    } else {
      message.delegationAllowed = false;
    }
    if (object.withdrawn !== undefined && object.withdrawn !== null) {
      message.withdrawn = object.withdrawn;
    } else {
      message.withdrawn = "";
    }
    return message;
  },
};

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
