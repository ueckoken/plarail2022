import Ajv, { JSONSchemaType } from "ajv";
import { Blocks } from "../types";

const ajv = new Ajv();

interface SignalTaikenMessage {
  blocks: Blocks;
}

const signalTaikenMessageSchema: JSONSchemaType<SignalTaikenMessage> = {
  type: "object",
  properties: {
    blocks: { type: "object", required: [] },
  },
  required: ["blocks"],
  additionalProperties: true,
};

export const validateSignalTaikenMessage = ajv.compile(
  signalTaikenMessageSchema
);
