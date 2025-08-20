// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

/// various log levels
#[derive(Clone, PartialEq, Eq, Debug)]
pub enum LogLevel {
    Info,
    Warning,
    Error,
}
impl LogLevel{
    pub fn as_str(&self) -> &'static str {
        match *self {
            LogLevel::Info => "INFO",
            LogLevel::Warning => "WARNING",
            LogLevel::Error => "ERROR",
        }
    }
}
/// primary function for emitting logs
pub fn log(level: LogLevel, message: &str) -> String {
     match level {
        LogLevel::Info => info(message),
        LogLevel::Warning => warn(message),
        LogLevel::Error => error(message),
     }
}

pub fn info(message: &str) -> String {
    format!("[{}]: {}",LogLevel::Info.as_str(), message)
}

pub fn warn(message: &str) -> String {
    format!("[{}]: {}",LogLevel::Warning.as_str(), message)
}

pub fn error(message: &str) -> String {
    format!("[{}]: {}",LogLevel::Error.as_str(), message)
}