START TRANSACTION;

CREATE TABLE `Country` (
  `ISO` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `Name` varchar(255) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`ISO`)
);

CREATE TABLE `Promo` (
  `Id` int(11) NOT NULL AUTO_INCREMENT,
  `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `UpdatedAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated by system, dont update manually.',
  `Name` varchar(256) NOT NULL COMMENT 'Show to user, e.g. "mega show promo"\nDoesnt have to be unique.',
  `Description` varchar(1000) NOT NULL COMMENT 'Internal, not shown to User.',
  `Tandc` varchar(4000) DEFAULT NULL COMMENT 'Optional terms and conditions.  Ideally we would force this to be plain text (to avoid injection) and replace bank lines with <p) and CR/LF with <br> etc.',
  `StartAt` timestamp NULL,
  `EndAt` timestamp NULL,
  `AmountPerAccount` decimal(12,6) unsigned NOT NULL COMMENT 'max number of sky as a decimal. E.g. 0.5 = half a sky',
  `MaxAccounts` int(11) NOT NULL COMMENT 'the max number of people who can take part, also the max number of generated accounts for this promo.',
  `EnabledYN` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'Enabled by default (as startAt will usually be in the future)',
  `ShowKeyYN` tinyint(1) NOT NULL DEFAULT '1' COMMENT 'The default is to show the Key in the UI after they register. IF this is switched off, then either email should be switched on, or it has to be done manually (insecure)',
  `EmailKeyYN` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'Not support for phase 1 - for future use. Dont put in UI until supported.',
  `AdminEmail` varchar(256) COLLATE utf8_unicode_ci NOT NULL COMMENT 'Who should be alerted when promo starts/finishes, or if there are errors etc.\nWhen we have operators, we could remove this.',
  `SourceKey` varchar(100) COLLATE utf8_unicode_ci NOT NULL COMMENT 'This is generated by the system when they sumbit the request to create the promotion. \nThe operator will then need to transfer amountPerAccount * maxAccounts into this account.',
  `CleanupKey` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'The operator will input this. Any sky left in the wallet when the promotion finishes (end date reached) will be transfered back to this account.\nThis field is optional.  If they dont enter it, then someone will have to manually clean up using the servers wallet.',
  `CleanedUpAmount` decimal(12,6) unsigned DEFAULT NULL COMMENT 'Will be null until the cleanup runs and transfers something back. \nIf nothing to transfer (all promos claimed) then this amount will be zero.',
  `CleanedUpAt` timestamp NULL DEFAULT NULL COMMENT 'Only set if cleanedUp when promo ends.  Will be null if no sky sent back to cleanup account.',
  PRIMARY KEY (`Id`)
);

CREATE TABLE `PromoCode` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` timestamp NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Auto value, dont insert this value.',
  `PromoId` int(11) NOT NULL,
  `Code` varchar(40) COLLATE utf8_unicode_ci NOT NULL COMMENT 'A system generated unique code (could even be an account public key)',
  PRIMARY KEY (`Id`),
  KEY `FK_promoCode_promo_id` (`PromoId`),
  CONSTRAINT `FK_promoCode_promo_id` FOREIGN KEY (`PromoId`) REFERENCES `Promo` (`Id`),
  UNIQUE (`PromoId` ,`Code`)
);


CREATE TABLE `Registration` (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'System generated, do not write',
  `UpdatedAt` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'This should never be set unless manual hack.  System generaged, do not write.',
  `PromoId` int(11) NOT NULL,
  `PromoCodeId` bigint(20) NOT NULL,
  `FirstName` varchar(100) DEFAULT NULL,
  `LastName` varchar(100) DEFAULT NULL,
  `Email` varchar(256) COLLATE utf8_unicode_ci NOT NULL,
  `Mobile` varchar(40) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'Must be ASCII only.  e.g. +44 777 1771717',
  `AddressLine1` varchar(256) DEFAULT NULL,
  `AddressLine2` varchar(256) DEFAULT NULL,
  `CountryCode` varchar(20) COLLATE utf8_unicode_ci NOT NULL,
  `State` varchar(256) DEFAULT NULL,
  `City` varchar(256) DEFAULT NULL,
  `Postcode` varchar(56) DEFAULT NULL,
  `IP` varchar(40) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'ip v4 or v6\nASCII only (no emojis!)',
  `UserAgent` varchar(1000) DEFAULT NULL COMMENT 'The data must be truncated if > 1000',
  `PublicKey` varchar(100) DEFAULT NULL COMMENT 'The public key of the account created for this user.  Private key is not stored in the DB. It is shown to the User directly (or emailed)',
  `Amount` decimal(12,6) DEFAULT NULL COMMENT 'The amount transfered to this account.',
  `Status` varchar(100) DEFAULT NULL COMMENT 'Status of code activation (completed, pending, rejected)',
  `RejectionCode` int(11) DEFAULT NULL COMMENT 'Rejection code if status is rejected. 100 - max accounts reached, 101 - invaid redeem code, 102 - duplicate, 103 - aborted ',
  `TransferError` varchar(1000) DEFAULT NULL COMMENT 'Hopefully Null most of the time.',
  `TransferTransactionId` varchar(100) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`Id`),
  KEY `FK_registration_promo_id` (`PromoId`),
  KEY `FK_registration_promo_code_id` (`PromoCodeId`),
  KEY `FK_registration_country_id` (`CountryCode`),
  CONSTRAINT `FK_registration_promo_id` FOREIGN KEY (`PromoId`) REFERENCES `Promo` (`Id`),
  CONSTRAINT `FK_registration_country_id` FOREIGN KEY (`CountryCode`) REFERENCES `Country` (`ISO`),
  CONSTRAINT `FK_registration_promo_code_id` FOREIGN KEY (`PromoCodeId`) REFERENCES `PromoCode` (`Id`)
);

COMMIT;
