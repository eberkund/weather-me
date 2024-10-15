import {
  Calendar,
  CloudFog,
  Cog,
  Compass,
  Eye,
  MapPin,
  Sun,
  X,
} from "lucide-react";
import { Link } from "wouter";
import logo from "~/assets/logo.svg";
import mapPin from "~/assets/mapPin.svg";
import { Button } from "~/components/Button";
import { Card } from "~/components/Card";
import cloudy from "~/assets/darksky/cloudy.svg";
import sunny from "~/assets/darksky/clear-day.svg";
import rainy from "~/assets/darksky/rain.svg";
import windy from "~/assets/darksky/wind.svg";
import {
  addPlace,
  Place,
  getPlaces,
  removeCity,
  getSettings,
} from "~/services/storage";
import { useEffect, useState } from "react";
import { PlaceKit } from "@placekit/autocomplete-react";
import "@placekit/autocomplete-js/dist/placekit-autocomplete.css";
import environment from "~/services/environment";
import { PKResult } from "@placekit/client-js";
import { client, providersMap } from "~/services/api";
import {
  GetCurrentResponse,
  GetForecastResponse,
} from "~/gen/weather/v1/weather_pb";
import dayjs from "dayjs";
import utc from "dayjs/plugin/utc";
import objectSupport from "dayjs/plugin/objectSupport";
import { twMerge } from "tailwind-merge";

dayjs.extend(utc);
dayjs.extend(objectSupport);

export function Index() {
  const [places, setPlaces] = useState<Place[]>(getPlaces());
  const [current, setCurrent] = useState<GetCurrentResponse | null>(null);
  const [forecast, setForecast] = useState<GetForecastResponse | null>(null);
  const [active, setActive] = useState(0);

  useEffect(() => {
    const fetchData = async () => {
      if (active >= places.length) {
        setCurrent(null);
        return;
      }
      (async () => {
        const current = await client.getCurrent({
          longitude: places[active].lon,
          latitude: places[active].lat,
          provider: providersMap.get(getSettings().provider),
        });
        setCurrent(current);
      })();
      (async () => {
        const forecast = await client.getForecast({
          longitude: places[active].lon,
          latitude: places[active].lat,
          provider: providersMap.get(getSettings().provider),
        });
        console.log(forecast);
        setForecast(forecast);
      })();
    };
    fetchData();
  }, [places, active]);

  function onRemoveCity(idx: number) {
    places.splice(idx, 1);
    setPlaces([...places]);
    removeCity(idx);
  }

  function onPick(value: string, item: PKResult, index: number) {
    console.log({ value, item, index });
    const place = {
      city: item.city,
      country: item.country,
      lat: item.lat!,
      lon: item.lng!,
    };
    addPlace(place);
    setPlaces([...places, place]);
  }

  function selectPlace(index: number) {
    setActive(index);
  }

  return (
    <div className="flex min-h-full">
      <div className="flex-initial min-w-96 bg-white/30 p-8">
        <img src={logo} alt="" />
        <div className="my-5">
          <PlaceKit
            geolocation={false}
            onPick={onPick}
            apiKey={environment.VITE_PLACEKIT_API_KEY}
          />
        </div>
        <div className="flex flex-col gap-y-5">
          {places.map((place, i) => (
            <Card
              key={i}
              className={twMerge(
                "hover:bg-slate-300 active:ring-2",
                active === i && "ring-2 ring-lime-300"
              )}
              onClick={() => selectPlace(i)}
            >
              <div className="flex items-center gap-x-3">
                <img src={mapPin} />
                <div className="flex-grow">
                  <h2 className="font-semibold">{place.city}</h2>
                  <p className="text-secondary text-sm">{place.country}</p>
                </div>
                <X onClick={() => onRemoveCity(i)} />
              </div>
            </Card>
          ))}
        </div>
      </div>
      <div className="w-full">
        <div className="container p-8">
          <div className="flex place-content-between">
            <div className="flex items-center gap-x-2 text-white">
              <MapPin size={24} />
              <h2 className="font-light">Washington DC, USA</h2>
            </div>
            <div>
              <Button asChild>
                <Link href="/settings">
                  <Cog size={24} />
                </Link>
              </Button>
            </div>
          </div>
          <div className="grid grid-flow-col grid-cols-3 gap-5">
            <Card variant="dark">
              <h2>Humidity</h2>
              <CloudFog size={72} className="mx-auto my-3" />
              {current?.humidity && `${current.humidity}%`}
            </Card>
            <Card variant="dark">
              <h2>Visibility</h2>
              <Eye size={72} className="mx-auto my-3" />
              {current?.visibility && `${current.visibility} km`}
            </Card>
            <Card variant="dark">
              <h2>UV Index</h2>
              <Sun size={72} className="mx-auto my-3" />
              <div className="relative rounded-full h-2 bg-gradient-to-r from-green-500 via-yellow-500 to-red-500 my-3">
                {current?.uvIndex !== undefined && (
                  <div
                    style={{ left: `${(current.uvIndex / 11) * 100}%` }}
                    className="h-2 w-2 absolute rounded-full bg-white shadow-sm"
                  />
                )}
              </div>
              {current?.uvIndex !== undefined && (
                <p>
                  {current.uvIndex}
                  <span className="text-sm ml-1">UV</span>
                </p>
              )}
            </Card>
            <Card variant="dark">
              <h2>Wind</h2>
              <Compass size={72} className="mx-auto my-3" />
              {current?.windSpeed !== undefined && `${current.windSpeed} km/h`}
            </Card>
            <div className="row-span-2 text-white text-center">
              <div className="flex flex-col gap-y-2">
                <div>{dayjs().format("h:mm A")}</div>
                <div>{dayjs().format("dddd, D MMMM YYYY")}</div>
              </div>
              {current?.temperature !== undefined && (
                <div className="text-7xl mt-5">
                  {`${current.temperature.toFixed()} °C`}
                </div>
              )}
            </div>
          </div>
          <div className="mt-8">
            <div className="flex items-center gap-x-2 text-white mb-5">
              <Calendar size={24} />
              <h2 className="font-light">7-day Forecast</h2>
            </div>
            <div className="grid grid-cols-1 md:grid-cols-3 xl:grid-cols-7 gap-3">
              {forecast?.days.map((day, i) => (
                <Card key={i} className="text-center">
                  <h3 className="font-medium mb-1">
                    {day.date &&
                      dayjs({
                        year: day.date.year,
                        month: day.date.month - 1,
                        day: day.date.day - 1,
                      }).format("ddd")}
                  </h3>
                  <hr className="divider" />
                  {day.condition}
                  <p className="text-white text-4xl">{day.temperature}°</p>
                </Card>
              ))}
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
