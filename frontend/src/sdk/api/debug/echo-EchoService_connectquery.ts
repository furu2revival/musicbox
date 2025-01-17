// @generated by protoc-gen-connect-query v0.4.1 with parameter "target=ts"
// @generated from file api/debug/echo.proto (package api.debug, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { createQueryService } from "@bufbuild/connect-query";
import { MethodKind } from "@bufbuild/protobuf";
import {
	EchoServiceEchoV1Request,
	EchoServiceEchoV1Response,
} from "./echo_pb.js";

export const typeName = "api.debug.EchoService";

/**
 * @generated from rpc api.debug.EchoService.EchoV1
 */
export const echoV1 = createQueryService({
	service: {
		methods: {
			echoV1: {
				name: "EchoV1",
				kind: MethodKind.Unary,
				I: EchoServiceEchoV1Request,
				O: EchoServiceEchoV1Response,
			},
		},
		typeName: "api.debug.EchoService",
	},
}).echoV1;
