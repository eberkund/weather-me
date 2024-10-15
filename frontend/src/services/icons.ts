import { Condition } from "~/gen/weather/v1/weather_pb";
import cloudy from "~/assets/darksky/cloudy.svg";
import rain from "~/assets/darksky/rain.svg";
import clearDay from "~/assets/darksky/clear-day.svg";
import partlyCloudyDay from "~/assets/darksky/partly-cloudy-day.svg";
import snow from "~/assets/darksky/snow.svg";
import fog from "~/assets/darksky/fog.svg";
import thunderstorm from "~/assets/darksky/thunderstorm.svg";
import question from "~/assets/question.svg";

export function conditionIcon(conditon: Condition) {
    switch (conditon) {
      case Condition.CLEAR:
        return clearDay;
      case Condition.CLOUDY:
        return cloudy;
      case Condition.PARTLY_CLOUDY:
        return partlyCloudyDay;
      case Condition.OVERCAST:
        return partlyCloudyDay;
      case Condition.MIST:
        return fog;
      case Condition.SUNNY:
        return clearDay;
      case Condition.RAINY:
        return rain;
      case Condition.SNOWY:
        return snow;
      case Condition.STORMY:
        return thunderstorm;
      case Condition.UNSPECIFIED:
      default:
        return question;
    }
}