import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "~/App.tsx";
import "~/index.css";
import { WeatherService } from "~/gen/weather/v1/weather_connect";
import { createGrpcWebTransport } from "@connectrpc/connect-web";
import { createPromiseClient } from "@connectrpc/connect";

const transport = createGrpcWebTransport({
  baseUrl: "http://localhost:8080",
  useBinaryFormat: true,
  
});

const client = createPromiseClient(WeatherService, transport);
client.getWeather({ city: "London" }).then((response) => {
  console.log(response.weather);
}).catch((error) => { 
  console.error("Error:", error);
});

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <App />
  </StrictMode>
);
