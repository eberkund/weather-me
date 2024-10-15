import { createPromiseClient } from "@connectrpc/connect";
import { createConnectTransport } from "@connectrpc/connect-web";
import { WeatherService } from "~/gen/weather/v1/weather_connect";
import environment from "~/services/environment";
import { Provider } from "~/services/storage";
import { WeatherProvider } from "~/gen/weather/v1/weather_pb";

export const client = createPromiseClient(
  WeatherService,
  createConnectTransport({
    baseUrl: environment.VITE_BACKEND_URL,
    useBinaryFormat: true,
  })
);

export const providersMap: Map<Provider, WeatherProvider> = new Map([
  [Provider.OpenWeather, WeatherProvider.OPENWEATHER],
  [Provider.WeatherAPI, WeatherProvider.WEATHERAPI],
]);