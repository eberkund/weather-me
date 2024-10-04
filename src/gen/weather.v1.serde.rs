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
        let mut struct_ser = serializer.serialize_struct("weather.v1.GetWeatherRequest", len)?;
        if !self.city.is_empty() {
            struct_ser.serialize_field("city", &self.city)?;
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
        ];

        #[allow(clippy::enum_variant_names)]
        enum GeneratedField {
            City,
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
                while let Some(k) = map_.next_key()? {
                    match k {
                        GeneratedField::City => {
                            if city__.is_some() {
                                return Err(serde::de::Error::duplicate_field("city"));
                            }
                            city__ = Some(map_.next_value()?);
                        }
                    }
                }
                Ok(GetWeatherRequest {
                    city: city__.unwrap_or_default(),
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
