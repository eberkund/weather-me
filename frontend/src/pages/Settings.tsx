import { X } from "lucide-react";
import { SubmitHandler, useForm } from "react-hook-form";
import { Link } from "wouter";
import { Button } from "~/components/Button";
import { Card } from "~/components/Card";
import {
  Distance,
  Frequency,
  getSettings,
  Pressure,
  Provider,
  saveSettings,
  Temperature,
  Time,
  WindSpeed,
  type UserSettings,
} from "~/storage";

export function Settings() {
  const {
    register,
    handleSubmit,
    // watch,
    formState: { errors },
  } = useForm<UserSettings>({
    defaultValues: getSettings(),
  });

  const onSubmit: SubmitHandler<UserSettings> = (data) => console.log(data);

  function saveData(event: React.FormEvent<HTMLFormElement>) {
    console.log("saving data");
    const formData = new FormData(event.currentTarget);
    const formEntries = Object.fromEntries(formData.entries());
    saveSettings(formEntries as UserSettings);
  }

  return (
    <div className="p-8">
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-2xl font-bold">Settings</h1>
        <div>
          <Button asChild>
            <Link href="/">
              <X size={24} width={24} height={24} />
            </Link>
          </Button>
        </div>
      </div>
      <form className="grid grid-cols-2 gap-5" onChange={saveData}>
        <div className="col-span-2">
          <Card>
            <h2 className="text-settings text-xl font-semibold mb-3">
              Provider
            </h2>
            <div>
              <label className="text-settings">Forecast Provider</label>
              <select
                {...register("provider")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(Provider).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
          </Card>
        </div>
        <Card>
          <h2 className="text-settings text-xl font-semibold mb-3">Unit</h2>
          <div className="flex flex-col gap-y-5">
            <div>
              <label className="text-settings">Temperature</label>
              <select
                {...register("temperature")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(Temperature).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
            <div>
              <label className="text-settings">Wind Speed</label>
              <select
                {...register("windSpeed")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(WindSpeed).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
            <div>
              <label className="text-settings">Pressure</label>
              <select
                {...register("pressure")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(Pressure).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
            <div>
              <label className="text-settings">Distance</label>
              <select
                {...register("distance")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(Distance).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
          </div>
        </Card>
        <Card>
          <h2 className="text-settings text-xl font-semibold mb-3">General</h2>
          <div className="flex flex-col gap-y-5">
            <div>
              <label className="text-settings">Notifications</label>
              <select
                {...register("notifications")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                <option>On</option>
                <option>Off</option>
              </select>
            </div>
            <div>
              <label className="text-settings">Frequency</label>
              <select
                {...register("frequency")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(Frequency).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
            <div>
              <label className="text-settings">Time</label>
              <select
                {...register("time")}
                className="mt-2 block w-full rounded-md border-0 py-1.5 pl-3 pr-10 text-gray-900  focus:ring-2 focus:ring-indigo-600 sm:text-sm sm:leading-6"
              >
                {Object.values(Time).map((value) => (
                  <option key={value}>{value}</option>
                ))}
              </select>
            </div>
          </div>
        </Card>
      </form>
    </div>
  );
}
