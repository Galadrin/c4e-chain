import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgStoreSignature } from "./types/c4e-chain/cfesignature/tx";
import { MsgPublishReferencePayloadLink } from "./types/c4e-chain/cfesignature/tx";
import { MsgCreateAccount } from "./types/c4e-chain/cfesignature/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/chain4energy.c4echain.cfesignature.MsgStoreSignature", MsgStoreSignature],
    ["/chain4energy.c4echain.cfesignature.MsgPublishReferencePayloadLink", MsgPublishReferencePayloadLink],
    ["/chain4energy.c4echain.cfesignature.MsgCreateAccount", MsgCreateAccount],
    
];

export { msgTypes }