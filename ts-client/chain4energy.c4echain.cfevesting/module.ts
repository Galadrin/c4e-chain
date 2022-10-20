// Generated by Ignite ignite.com/cli

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient, DeliverTxResponse } from "@cosmjs/stargate";
import { EncodeObject, GeneratedType, OfflineSigner, Registry } from "@cosmjs/proto-signing";
import { msgTypes } from './registry';
import { IgniteClient } from "../client"
import { MissingWalletError } from "../helpers"
import { Api } from "./rest";
import { MsgWithdrawAllAvailable } from "./types/cfevesting/tx";
import { MsgSendToVestingAccount } from "./types/cfevesting/tx";
import { MsgCreateVestingPool } from "./types/cfevesting/tx";
import { MsgCreateVestingAccount } from "./types/cfevesting/tx";


export { MsgWithdrawAllAvailable, MsgSendToVestingAccount, MsgCreateVestingPool, MsgCreateVestingAccount };

type sendMsgWithdrawAllAvailableParams = {
  value: MsgWithdrawAllAvailable,
  fee?: StdFee,
  memo?: string
};

type sendMsgSendToVestingAccountParams = {
  value: MsgSendToVestingAccount,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateVestingPoolParams = {
  value: MsgCreateVestingPool,
  fee?: StdFee,
  memo?: string
};

type sendMsgCreateVestingAccountParams = {
  value: MsgCreateVestingAccount,
  fee?: StdFee,
  memo?: string
};


type msgWithdrawAllAvailableParams = {
  value: MsgWithdrawAllAvailable,
};

type msgSendToVestingAccountParams = {
  value: MsgSendToVestingAccount,
};

type msgCreateVestingPoolParams = {
  value: MsgCreateVestingPool,
};

type msgCreateVestingAccountParams = {
  value: MsgCreateVestingAccount,
};


export const registry = new Registry(msgTypes);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
	prefix: string
	signer?: OfflineSigner
}

export const txClient = ({ signer, prefix, addr }: TxClientOptions = { addr: "http://localhost:26657", prefix: "cosmos" }) => {

  return {
		
		async sendMsgWithdrawAllAvailable({ value, fee, memo }: sendMsgWithdrawAllAvailableParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgWithdrawAllAvailable: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgWithdrawAllAvailable({ value: MsgWithdrawAllAvailable.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgWithdrawAllAvailable: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgSendToVestingAccount({ value, fee, memo }: sendMsgSendToVestingAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgSendToVestingAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgSendToVestingAccount({ value: MsgSendToVestingAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgSendToVestingAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateVestingPool({ value, fee, memo }: sendMsgCreateVestingPoolParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateVestingPool: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateVestingPool({ value: MsgCreateVestingPool.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateVestingPool: Could not broadcast Tx: '+ e.message)
			}
		},
		
		async sendMsgCreateVestingAccount({ value, fee, memo }: sendMsgCreateVestingAccountParams): Promise<DeliverTxResponse> {
			if (!signer) {
					throw new Error('TxClient:sendMsgCreateVestingAccount: Unable to sign Tx. Signer is not present.')
			}
			try {			
				const { address } = (await signer.getAccounts())[0]; 
				const signingClient = await SigningStargateClient.connectWithSigner(addr,signer,{registry, prefix});
				let msg = this.msgCreateVestingAccount({ value: MsgCreateVestingAccount.fromPartial(value) })
				return await signingClient.signAndBroadcast(address, [msg], fee ? fee : defaultFee, memo)
			} catch (e: any) {
				throw new Error('TxClient:sendMsgCreateVestingAccount: Could not broadcast Tx: '+ e.message)
			}
		},
		
		
		msgWithdrawAllAvailable({ value }: msgWithdrawAllAvailableParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfevesting.MsgWithdrawAllAvailable", value: MsgWithdrawAllAvailable.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgWithdrawAllAvailable: Could not create message: ' + e.message)
			}
		},
		
		msgSendToVestingAccount({ value }: msgSendToVestingAccountParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfevesting.MsgSendToVestingAccount", value: MsgSendToVestingAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgSendToVestingAccount: Could not create message: ' + e.message)
			}
		},
		
		msgCreateVestingPool({ value }: msgCreateVestingPoolParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfevesting.MsgCreateVestingPool", value: MsgCreateVestingPool.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateVestingPool: Could not create message: ' + e.message)
			}
		},
		
		msgCreateVestingAccount({ value }: msgCreateVestingAccountParams): EncodeObject {
			try {
				return { typeUrl: "/chain4energy.c4echain.cfevesting.MsgCreateVestingAccount", value: MsgCreateVestingAccount.fromPartial( value ) }  
			} catch (e: any) {
				throw new Error('TxClient:MsgCreateVestingAccount: Could not create message: ' + e.message)
			}
		},
		
	}
};

interface QueryClientOptions {
  addr: string
}

export const queryClient = ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

class SDKModule {
	public query: ReturnType<typeof queryClient>;
	public tx: ReturnType<typeof txClient>;
	
	public registry: Array<[string, GeneratedType]>;

	constructor(client: IgniteClient) {		
	
		this.query = queryClient({ addr: client.env.apiURL });
		this.tx = txClient({ signer: client.signer, addr: client.env.rpcURL, prefix: client.env.prefix ?? "cosmos" });
	}
};

const Module = (test: IgniteClient) => {
	return {
		module: {
			Chain4EnergyC4EchainCfevesting: new SDKModule(test)
		},
		registry: msgTypes
  }
}
export default Module;