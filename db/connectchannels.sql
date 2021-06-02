CREATE TABLE `connectchannels` (
  `Name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Platform` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Platform_ID` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Connect` char(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Announce` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;


INSERT INTO `connectchannels` (`Name`, `Platform`, `Platform_ID`, `Connect`, `Announce`) VALUES
('nouryqt', 'Twitch', '31437432 ', 'true', 'true'),
('nourybot', 'Twitch', '596581605 ', 'true', 'false');