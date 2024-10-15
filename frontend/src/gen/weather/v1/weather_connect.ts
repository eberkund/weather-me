// @generated by protoc-gen-connect-es v1.6.0 with parameter "target=ts"
// @generated from file weather/v1/weather.proto (package weather.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { GetCurrentRequest, GetCurrentResponse, GetForecastRequest, GetForecastResponse } from "./weather_pb.js";
import { MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service weather.v1.WeatherService
 */
export const WeatherService = {
  typeName: "weather.v1.WeatherService",
  methods: {
    /**
     * @generated from rpc weather.v1.WeatherService.GetCurrent
     */
    getCurrent: {
      name: "GetCurrent",
      I: GetCurrentRequest,
      O: GetCurrentResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc weather.v1.WeatherService.GetForecast
     */
    getForecast: {
      name: "GetForecast",
      I: GetForecastRequest,
      O: GetForecastResponse,
      kind: MethodKind.Unary,
    },
  }
} as const;

