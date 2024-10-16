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

export interface Place {
  city: string;
  country: string;
  lat: number;
  lon: number;
}

export function addPlace(place: Place) {
  const places = getPlaces();
  localStorage.setItem("places", JSON.stringify([...places, place]));
}

export function removeCity(idx: number) {
  const places = getPlaces();
  places.splice(idx, 1);
  localStorage.setItem("places", JSON.stringify(places));
}

export function getPlaces(): Place[] {
  const places = localStorage.getItem("places");
  return places ? JSON.parse(places) : [];
}
