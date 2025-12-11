package lib

// Note: With the removal of filesMaxKeepTime.json, files are now stored indefinitely
// until they are manually deleted or the server reaches max storage capacity.
// Automatic cleanup based on time is no longer supported.
// If space is needed, older files should be manually deleted via the DELETE /files/:file endpoint.
