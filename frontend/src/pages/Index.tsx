import { Calendar, MapPin, X } from "lucide-react";
import { Link } from "wouter";
import cog from "~/assets/cog.svg";
import logo from "~/assets/logo.svg";
import mapPin from "~/assets/mapPin.svg";
import { Button } from "~/components/Button";
import { Card } from "~/components/Card";
import cloudy from "~/assets/darksky/cloudy.svg";
import sunny from "~/assets/darksky/clear-day.svg";
import rainy from "~/assets/darksky/rain.svg";
import windy from "~/assets/darksky/wind.svg";
import { addCity, City } from "~/storage";
import { useState } from "react";

const defaultCities = [
  {
    name: "Libson",
    country: "Portugal",
  },
  {
    name: "Kyiv",
    country: "Ukraine",
  },
  {
    name: "Berlin",
    country: "Germany",
  },
];

const days = [
  {
    day: "Monday",
    temperature: 20,
    weather: rainy,
  },
  {
    day: "Tuesday",
    temperature: 22,
    weather: rainy,
  },
  {
    day: "Wednesday",
    temperature: 25,
    weather: sunny,
  },
  {
    day: "Thursday",
    temperature: 24,
    weather: sunny,
  },
  {
    day: "Friday",
    temperature: 26,
    weather: cloudy,
  },
  {
    day: "Saturday",
    temperature: 27,
    weather: sunny,
  },
  {
    day: "Sunday",
    temperature: 28,
    weather: windy,
  },
];

export function Index() {
  const [cities, setCities] = useState<City[]>(defaultCities);

  function onCitySubmit(event: React.FormEvent<HTMLFormElement>) {
    event.preventDefault();
    const form = event.currentTarget;
    const city = form.elements.namedItem("city") as HTMLInputElement;
    addCity({ name: city.value });
  }

  function removeCity(idx: number) {
    cities.splice(idx, 1);
    setCities([...cities]);
  }

  return (
    <div className="flex h-full">
      <div className="flex-initial min-w-96 bg-white/30 p-8">
        <img src={logo} alt="" />
        <div className="my-5">
          <form onSubmit={onCitySubmit}>
            <input
              name="city"
              type="text"
              placeholder="City"
              className="block w-full rounded-md border-0 py-1.5 text-gray-900 shadow-sm placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600"
            />
          </form>
        </div>
        <div className="flex flex-col gap-y-5">
          {cities.map((city, i) => (
            <Card key={i}>
              <div className="flex items-center gap-x-3">
                <img src={mapPin} />
                <div className="flex-grow">
                  <h2 className="font-semibold">{city.name}</h2>
                  <p className="text-secondary text-sm">{city.country}</p>
                </div>
                <X onClick={() => removeCity(i)} />
              </div>
            </Card>
          ))}
        </div>
      </div>
      <div className="w-full">
        <div className="container p-5">
          <div className="flex place-content-between">
            <div className="flex items-center gap-x-2 text-white">
              <MapPin size={24} />
              <h2 className="font-light">Washington DC, USA</h2>
            </div>
            <div>
              <Button asChild>
                <Link href="/settings">
                  <img src={cog} alt="Settings" />
                </Link>
              </Button>
            </div>
          </div>
          <div className="grid grid-flow-col grid-cols-3 gap-5">
            <Card variant="dark">
              <h2>Wind Status</h2>
            </Card>
            <Card variant="dark">
              <h2>Wind Status</h2>
            </Card>
            <Card variant="dark">
              <h2>Wind Status</h2>
            </Card>
            <Card variant="dark">
              <h2>Wind Status</h2>
            </Card>
            <div className="row-span-2">
              <h2>Today</h2>
            </div>
          </div>
          <div>
            <div className="flex items-center gap-x-2 text-white mb-5">
              <Calendar size={24} />
              <h2 className="font-light">7-day Forecast</h2>
            </div>
            <div className="grid grid-cols-7 gap-3">
              {days.map((day) => (
                <Card key={day.day} className="text-center">
                  <h3 className="font-medium mb-1">{day.day}</h3>
                  <hr className="divider" />
                  <img src={day.weather} alt="" className="my-2" />
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
