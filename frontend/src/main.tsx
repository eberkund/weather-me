import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "~/App.tsx";
import "~/index.css";
import "@fontsource/poppins/300.css";
import "@fontsource/poppins/400.css";
import "@fontsource/poppins/500.css";
import "@fontsource/poppins/600.css";
import "@fontsource/poppins/700.css";
import "@fontsource/poppins/800.css";
import "@fontsource/poppins/900.css";
import { WeatherService } from "~/gen/weather/v1/weather_connect";
import { createGrpcWebTransport } from "@connectrpc/connect-web";
import { createPromiseClient } from "@connectrpc/connect";
import { WeatherProvider } from "./gen/weather/v1/weather_pb";

const client = createPromiseClient(
  WeatherService,
  createGrpcWebTransport({
    baseUrl: "http://localhost:8080",
    useBinaryFormat: true,
  })
);
const response = await client.getWeather({
  city: "London",
  provider: WeatherProvider.OPEN_WEATHER,
});
console.log(response.weather);

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
  </StrictMode>
);
