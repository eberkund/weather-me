export interface UserSettings {
  provider: Provider;
  temperature: Temperature;
  windSpeed: string;
  pressure: string;
  distance: string;
  notifications: boolean;
  frequency: string;
  time: string;
}

export enum Provider {
  OpenWeather = "OpenWeather",
  WeatherAPI = "WeatherAPI",
}

export enum Temperature {
  Celsius = "Celsius",
  Farenheit = "Farenheit",
}

export enum WindSpeed {
  MetersPerSecond = "m/s",
  KilometersPerHour = "km/hr",
  FeetPerSecond = "ft/s",
  MilesPerHour = "mph",
  Knots = "knots",
}

export enum Pressure {
  Hectopascals = "hPa",
  Millibars = "mb",
}

export enum Distance {
  Kilometers = "km",
  Miles = "mi",
}

export enum Frequency {
  EveryMinute = "every minute",
  FifteenMinutes = "15 minutes",
  ThirtyMinutes = "30 minutes",
  Hourly = "hourly",
}

export enum Time {
  Clock12Hour = "12 hour",
  Clock24Hour = "24 hour",
}

const defaultSettings: UserSettings = {
  provider: Provider.OpenWeather,
  temperature: Temperature.Celsius,
  windSpeed: WindSpeed.KilometersPerHour,
  pressure: Pressure.Hectopascals,
  distance: Distance.Kilometers,
  notifications: false,
  frequency: Frequency.FifteenMinutes,
  time: Time.Clock12Hour,
};

export function saveSettings(settings: UserSettings) {
  localStorage.setItem("settings", JSON.stringify(settings));
}

export function getSettings(): UserSettings {
  const settings = localStorage.getItem("settings");
  return settings ? JSON.parse(settings) : defaultSettings;
}

export interface City {
  name: string;
}

export function addCity(city: City) {
  const cities = getCities();
  localStorage.setItem("cities", JSON.stringify([...cities, city]));
}

export function removeCity(idx: number) {
  const cities = getCities();
  cities.splice(idx, 1);
  localStorage.setItem("cities", JSON.stringify(cities));
}

export function getCities(): City[] {
  const cities = localStorage.getItem("cities");
  return cities ? JSON.parse(cities) : [];
}
