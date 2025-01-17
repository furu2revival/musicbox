// @generated by protoc-gen-es v1.10.0 with parameter "target=ts"
// @generated from file custom_option/custom_option.proto (package custom_option, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type {
	BinaryReadOptions,
	FieldList,
	JsonReadOptions,
	JsonValue,
	PartialMessage,
	PlainMessage,
} from "@bufbuild/protobuf";
import { Message, MethodOptions, proto3 } from "@bufbuild/protobuf";
import {
	ErrorCode_Method,
	ErrorSeverity,
} from "../api/api_errors/api_errors_pb.js";

/**
 * @generated from message custom_option.MethodErrorDefinition
 */
export class MethodErrorDefinition extends Message<MethodErrorDefinition> {
	/**
	 * @generated from field: api.api_errors.ErrorCode.Method code = 1;
	 */
	code = ErrorCode_Method.UNSPECIFIED;

	/**
	 * @generated from field: api.api_errors.ErrorSeverity severity = 2;
	 */
	severity = ErrorSeverity.UNSPECIFIED;

	/**
	 * @generated from field: string message = 3;
	 */
	message = "";

	constructor(data?: PartialMessage<MethodErrorDefinition>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "custom_option.MethodErrorDefinition";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{
			no: 1,
			name: "code",
			kind: "enum",
			T: proto3.getEnumType(ErrorCode_Method),
		},
		{
			no: 2,
			name: "severity",
			kind: "enum",
			T: proto3.getEnumType(ErrorSeverity),
		},
		{ no: 3, name: "message", kind: "scalar", T: 9 /* ScalarType.STRING */ },
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): MethodErrorDefinition {
		return new MethodErrorDefinition().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): MethodErrorDefinition {
		return new MethodErrorDefinition().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): MethodErrorDefinition {
		return new MethodErrorDefinition().fromJsonString(jsonString, options);
	}

	static equals(
		a: MethodErrorDefinition | PlainMessage<MethodErrorDefinition> | undefined,
		b: MethodErrorDefinition | PlainMessage<MethodErrorDefinition> | undefined
	): boolean {
		return proto3.util.equals(MethodErrorDefinition, a, b);
	}
}

/**
 * @generated from message custom_option.MethodOption
 */
export class MethodOption extends Message<MethodOption> {
	/**
	 * @generated from field: repeated custom_option.MethodErrorDefinition method_error_definitions = 1;
	 */
	methodErrorDefinitions: MethodErrorDefinition[] = [];

	constructor(data?: PartialMessage<MethodOption>) {
		super();
		proto3.util.initPartial(data, this);
	}

	static readonly runtime: typeof proto3 = proto3;
	static readonly typeName = "custom_option.MethodOption";
	static readonly fields: FieldList = proto3.util.newFieldList(() => [
		{
			no: 1,
			name: "method_error_definitions",
			kind: "message",
			T: MethodErrorDefinition,
			repeated: true,
		},
	]);

	static fromBinary(
		bytes: Uint8Array,
		options?: Partial<BinaryReadOptions>
	): MethodOption {
		return new MethodOption().fromBinary(bytes, options);
	}

	static fromJson(
		jsonValue: JsonValue,
		options?: Partial<JsonReadOptions>
	): MethodOption {
		return new MethodOption().fromJson(jsonValue, options);
	}

	static fromJsonString(
		jsonString: string,
		options?: Partial<JsonReadOptions>
	): MethodOption {
		return new MethodOption().fromJsonString(jsonString, options);
	}

	static equals(
		a: MethodOption | PlainMessage<MethodOption> | undefined,
		b: MethodOption | PlainMessage<MethodOption> | undefined
	): boolean {
		return proto3.util.equals(MethodOption, a, b);
	}
}

/**
 * @generated from extension: custom_option.MethodOption method_option = 50000;
 */
export const method_option = proto3.makeExtension<MethodOptions, MethodOption>(
	"custom_option.method_option",
	MethodOptions,
	() => ({ no: 50000, kind: "message", T: MethodOption })
);
