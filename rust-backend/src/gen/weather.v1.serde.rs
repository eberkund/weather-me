// @generated
impl serde::Serialize for GetWeatherRequest {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.city.is_empty() {
            len += 1;
        }
        if self.provider != 0 {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetWeatherRequest", len)?;
        if !self.city.is_empty() {
            struct_ser.serialize_field("city", &self.city)?;
        }
        if self.provider != 0 {
            let v = WeatherProvider::try_from(self.provider)
                .map_err(|_| serde::ser::Error::custom(format!("Invalid variant {}", self.provider)))?;
            struct_ser.serialize_field("provider", &v)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetWeatherRequest {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "city",
            "provider",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            City,
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
                            "city" => Ok(GeneratedField::City),
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
            type Value = GetWeatherRequest;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.GetWeatherRequest")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetWeatherRequest, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut city__ = None;
                let mut provider__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::City => {
                            if city__.is_some() {
                                return Err(serde::de::Error::duplicate_field("city"));
                            }
                            city__ = Some(map_.next_value()?);
                        }
                        GeneratedField::Provider => {
                            if provider__.is_some() {
                                return Err(serde::de::Error::duplicate_field("provider"));
                            }
                            provider__ = Some(map_.next_value::<WeatherProvider>()? as i32);
                        }
                    }
                }
                Ok(GetWeatherRequest {
                    city: city__.unwrap_or_default(),
                    provider: provider__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.GetWeatherRequest", FIELDS, GeneratedVisitor)
    }
}
impl serde::Serialize for GetWeatherResponse {
    #[allow(deprecated)]
    fn serialize<S>(&self, serializer: S) -> std::result::Result<S::Ok, S::Error>
    where
        S: serde::Serializer,
    {
        use serde::ser::SerializeStruct;
        let mut len = 0;
        if !self.weather.is_empty() {
            len += 1;
        }
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetWeatherResponse", len)?;
        if !self.weather.is_empty() {
            struct_ser.serialize_field("weather", &self.weather)?;
        }
        struct_ser.end()
    }
}
impl<'de> serde::Deserialize<'de> for GetWeatherResponse {
    #[allow(deprecated)]
    fn deserialize<D>(deserializer: D) -> std::result::Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        const FIELDS: &[&str] = &[
            "weather",
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            Weather,
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
                            "weather" => Ok(GeneratedField::Weather),
                            _ => Err(serde::de::Error::unknown_field(value, FIELDS)),
                        }
                    }
                }
                deserializer.deserialize_identifier(GeneratedVisitor)
            }
        }
        struct GeneratedVisitor;
        impl<'de> serde::de::Visitor<'de> for GeneratedVisitor {
            type Value = GetWeatherResponse;

            fn expecting(&self, formatter: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
                formatter.write_str("struct weather.v1.GetWeatherResponse")
            }

            fn visit_map<V>(self, mut map_: V) -> std::result::Result<GetWeatherResponse, V::Error>
                where
                    V: serde::de::MapAccess<'de>,
            {
                let mut weather__ = None;
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::Weather => {
                            if weather__.is_some() {
                                return Err(serde::de::Error::duplicate_field("weather"));
                            }
                            weather__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(GetWeatherResponse {
                    weather: weather__.unwrap_or_default(),
                })
            }
        }
        deserializer.deserialize_struct("weather.v1.GetWeatherResponse", FIELDS, GeneratedVisitor)
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
            Self::OpenWeather => "WEATHER_PROVIDER_OPEN_WEATHER",
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
            "WEATHER_PROVIDER_OPEN_WEATHER",
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
                    "WEATHER_PROVIDER_OPEN_WEATHER" => Ok(WeatherProvider::OpenWeather),
                    _ => Err(serde::de::Error::unknown_variant(value, FIELDS)),
                }
            }
        }
        deserializer.deserialize_any(GeneratedVisitor)
    }
}
