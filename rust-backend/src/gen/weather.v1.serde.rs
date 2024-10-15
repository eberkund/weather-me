// @generated
impl serde::Serialize for Condition {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "CONDITION_UNSPECIFIED",
            Self::Clear => "CLEAR",
            Self::Cloudy => "CLOUDY",
            Self::PartlyCloudy => "PARTLY_CLOUDY",
            Self::Overcast => "OVERCAST",
            Self::Mist => "MIST",
            Self::Sunny => "SUNNY",
            Self::Rainy => "RAINY",
            Self::Snowy => "SNOWY",
            Self::Stormy => "STORMY",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for Condition {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "CONDITION_UNSPECIFIED",
            "CLEAR",
            "CLOUDY",
            "PARTLY_CLOUDY",
            "OVERCAST",
            "MIST",
            "SUNNY",
            "RAINY",
            "SNOWY",
            "STORMY",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = Condition;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "CONDITION_UNSPECIFIED" => Ok(Condition::Unspecified),
                    "CLEAR" => Ok(Condition::Clear),
                    "CLOUDY" => Ok(Condition::Cloudy),
                    "PARTLY_CLOUDY" => Ok(Condition::PartlyCloudy),
                    "OVERCAST" => Ok(Condition::Overcast),
                    "MIST" => Ok(Condition::Mist),
                    "SUNNY" => Ok(Condition::Sunny),
                    "RAINY" => Ok(Condition::Rainy),
                    "SNOWY" => Ok(Condition::Snowy),
                    "STORMY" => Ok(Condition::Stormy),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
impl serde::Serialize for DailyForecast {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.temperature != 0. {
            len += 1;
        }
        if self.date.is_some() {
            len += 1;
        }
        if self.condition != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.DailyForecast", len)?;
        if self.temperature != 0. {
            struct_ser.serialize_field("temperature", &self.temperature)?;
        }
        if let Some(v) = self.date.as_ref() {
            struct_ser.serialize_field("date", v)?;
        }
        if self.condition != 0 {
            let v = Condition::try_from(self.condition)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.condition)))?;
            struct_ser.serialize_field("condition", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for DailyForecast {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "temperature",
            "date",
            "condition",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Temperature,
            Date,
            Condition,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "temperature" => Ok(GeneratedField::Temperature),
                            "date" => Ok(GeneratedField::Date),
                            "condition" => Ok(GeneratedField::Condition),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = DailyForecast;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.DailyForecast")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<DailyForecast, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut temperature__ = None;
                let mut date__ = None;
                let mut condition__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Temperature => {
                            if temperature__.is_some() {
                                return Err(serde::de::Error::duplicate_field("temperature"));
                            }
                            temperature__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Date => {
                            if date__.is_some() {
                                return Err(serde::de::Error::duplicate_field("date"));
                            }
                            date__ = map_.next_value()?;
                        }
                        GeneratedField::Condition => {
                            if condition__.is_some() {
                                return Err(serde::de::Error::duplicate_field("condition"));
                            }
                            condition__ = Some(map_.next_value::<Condition>()? as i32);
                        }
                    }
                }
                Ok(DailyForecast {
                    temperature: temperature__.unwrap_or_default(),
                    date: date__,
                    condition: condition__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.DailyForecast", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetCurrentRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.latitude != 0. {
            len += 1;
        }
        if self.longitude != 0. {
            len += 1;
        }
        if self.provider != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetCurrentRequest", len)?;
        if self.latitude != 0. {
            struct_ser.serialize_field("latitude", &self.latitude)?;
        }
        if self.longitude != 0. {
            struct_ser.serialize_field("longitude", &self.longitude)?;
        }
        if self.provider != 0 {
            let v = WeatherProvider::try_from(self.provider)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.provider)))?;
            struct_ser.serialize_field("provider", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetCurrentRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "latitude",
            "longitude",
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Latitude,
            Longitude,
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "latitude" => Ok(GeneratedField::Latitude),
                            "longitude" => Ok(GeneratedField::Longitude),
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GetCurrentRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.GetCurrentRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetCurrentRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut latitude__ = None;
                let mut longitude__ = None;
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Latitude => {
                            if latitude__.is_some() {
                                return Err(serde::de::Error::duplicate_field("latitude"));
                            }
                            latitude__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Longitude => {
                            if longitude__.is_some() {
                                return Err(serde::de::Error::duplicate_field("longitude"));
                            }
                            longitude__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value::<WeatherProvider>()? as i32);
                        }
                    }
                }
                Ok(GetCurrentRequest {
                    latitude: latitude__.unwrap_or_default(),
                    longitude: longitude__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.GetCurrentRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetCurrentResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.temperature != 0. {
            len += 1;
        }
        if self.humidity != 0. {
            len += 1;
        }
        if self.uv_index != 0. {
            len += 1;
        }
        if self.visibility != 0 {
            len += 1;
        }
        if self.wind_speed != 0. {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetCurrentResponse", len)?;
        if self.temperature != 0. {
            struct_ser.serialize_field("temperature", &self.temperature)?;
        }
        if self.humidity != 0. {
            struct_ser.serialize_field("humidity", &self.humidity)?;
        }
        if self.uv_index != 0. {
            struct_ser.serialize_field("uvIndex", &self.uv_index)?;
        }
        if self.visibility != 0 {
            struct_ser.serialize_field("visibility", &self.visibility)?;
        }
        if self.wind_speed != 0. {
            struct_ser.serialize_field("windSpeed", &self.wind_speed)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetCurrentResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "temperature",
            "humidity",
            "uv_index",
            "uvIndex",
            "visibility",
            "wind_speed",
            "windSpeed",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Temperature,
            Humidity,
            UvIndex,
            Visibility,
            WindSpeed,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "temperature" => Ok(GeneratedField::Temperature),
                            "humidity" => Ok(GeneratedField::Humidity),
                            "uvIndex" | "uv_index" => Ok(GeneratedField::UvIndex),
                            "visibility" => Ok(GeneratedField::Visibility),
                            "windSpeed" | "wind_speed" => Ok(GeneratedField::WindSpeed),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GetCurrentResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.GetCurrentResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetCurrentResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut temperature__ = None;
                let mut humidity__ = None;
                let mut uv_index__ = None;
                let mut visibility__ = None;
                let mut wind_speed__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Temperature => {
                            if temperature__.is_some() {
                                return Err(serde::de::Error::duplicate_field("temperature"));
                            }
                            temperature__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Humidity => {
                            if humidity__.is_some() {
                                return Err(serde::de::Error::duplicate_field("humidity"));
                            }
                            humidity__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::UvIndex => {
                            if uv_index__.is_some() {
                                return Err(serde::de::Error::duplicate_field("uvIndex"));
                            }
                            uv_index__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Visibility => {
                            if visibility__.is_some() {
                                return Err(serde::de::Error::duplicate_field("visibility"));
                            }
                            visibility__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::WindSpeed => {
                            if wind_speed__.is_some() {
                                return Err(serde::de::Error::duplicate_field("windSpeed"));
                            }
                            wind_speed__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                    }
                }
                Ok(GetCurrentResponse {
                    temperature: temperature__.unwrap_or_default(),
                    humidity: humidity__.unwrap_or_default(),
                    uv_index: uv_index__.unwrap_or_default(),
                    visibility: visibility__.unwrap_or_default(),
                    wind_speed: wind_speed__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.GetCurrentResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetForecastRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if self.latitude != 0. {
            len += 1;
        }
        if self.longitude != 0. {
            len += 1;
        }
        if self.provider != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetForecastRequest", len)?;
        if self.latitude != 0. {
            struct_ser.serialize_field("latitude", &self.latitude)?;
        }
        if self.longitude != 0. {
            struct_ser.serialize_field("longitude", &self.longitude)?;
        }
        if self.provider != 0 {
            let v = WeatherProvider::try_from(self.provider)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.provider)))?;
            struct_ser.serialize_field("provider", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetForecastRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "latitude",
            "longitude",
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Latitude,
            Longitude,
            Provider,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "latitude" => Ok(GeneratedField::Latitude),
                            "longitude" => Ok(GeneratedField::Longitude),
                            "provider" => Ok(GeneratedField::Provider),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GetForecastRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.GetForecastRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetForecastRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut latitude__ = None;
                let mut longitude__ = None;
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Latitude => {
                            if latitude__.is_some() {
                                return Err(serde::de::Error::duplicate_field("latitude"));
                            }
                            latitude__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Longitude => {
                            if longitude__.is_some() {
                                return Err(serde::de::Error::duplicate_field("longitude"));
                            }
                            longitude__ = 
                                Some(map_.next_value::<::pbjson::private::NumberDeserialize<_>>()?.0)
                            ;
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value::<WeatherProvider>()? as i32);
                        }
                    }
                }
                Ok(GetForecastRequest {
                    latitude: latitude__.unwrap_or_default(),
                    longitude: longitude__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.GetForecastRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetForecastResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.days.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetForecastResponse", len)?;
        if !self.days.is_empty() {
            struct_ser.serialize_field("days", &self.days)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetForecastResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "days",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Days,
        }
        impl<'de> serde::Deserialize<'de> for GeneratedField {
            fn deserialize<D>(deserializer: D) -> std::result::Result<GeneratedField, D::Error>
            where
                D: serde::Deserializer<'de>,
            {
                struct GeneratedVisitor;

                impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
                    type Value = GeneratedField;

                    fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                        write!(formatter, "expected one of: {:?}", &FIELDS)
                    }

                    #[allow(unused_variables)]
                    fn visit_str<E>(self, value: &str) -> std::result::Result<GeneratedField, E>
                    where
                        E: serde::de::Error,
                    {
                        match value {
                            "days" => Ok(GeneratedField::Days),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GetForecastResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.GetForecastResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetForecastResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut days__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Days => {
                            if days__.is_some() {
                                return Err(serde::de::Error::duplicate_field("days"));
                            }
                            days__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(GetForecastResponse {
                    days: days__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.GetForecastResponse", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for WeatherProvider {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        let variant = match self {
            Self::Unspecified => "WEATHER_PROVIDER_UNSPECIFIED",
            Self::Openweather => "WEATHER_PROVIDER_OPENWEATHER",
            Self::Weatherapi => "WEATHER_PROVIDER_WEATHERAPI",
        };
        serializer.serialize_str(variant)
    }
}
impl<'de> serde::Deserialize<'de> for WeatherProvider {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "WEATHER_PROVIDER_UNSPECIFIED",
            "WEATHER_PROVIDER_OPENWEATHER",
            "WEATHER_PROVIDER_WEATHERAPI",
        ];

        struct GeneratedVisitor;

        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = WeatherProvider;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                write!(formatter, "expected one of: {:?}", &FIELDS)
            }

            fn visit_i64<E>(self, v: i64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Signed(v), &self)
                    })
            }

            fn visit_u64<E>(self, v: u64) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                i32::try_from(v)
                    .ok()
                    .and_then(|x| x.try_into().ok())
                    .ok_or_else(|| {
                        serde::de::Error::invalid_value(serde::de::Unexpected::Unsigned(v), &self)
                    })
            }

            fn visit_str<E>(self, value: &str) -> std::result::Result<Self::Value, E>
            where
                E: serde::de::Error,
            {
                match value {
                    "WEATHER_PROVIDER_UNSPECIFIED" => Ok(WeatherProvider::Unspecified),
                    "WEATHER_PROVIDER_OPENWEATHER" => Ok(WeatherProvider::Openweather),
                    "WEATHER_PROVIDER_WEATHERAPI" => Ok(WeatherProvider::Weatherapi),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
