CREATE TABLE IF NOT EXISTS `connectchannels` (
  `Name` varchar(50) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Platform` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Connect` char(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `Announce` varchar(10) COLLATE utf8mb4_unicode_ci DEFAULT NULL
) ENGINE=MyISAM DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Name is Twitch username.
-- Platform is Twitch.
-- Connect is if we should join.
-- Announce is if we should say something when we join.
INSERT INTO `connectchannels` (`Name`, `Platform`, `Connect`, `Announce`) VALUES
('nouryqt', 'Twitch', , 'true', 'true'),
('nourybot', 'Twitch', , 'true', 'false');
COMMIT;