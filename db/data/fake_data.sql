INSERT INTO `Promo` (`Name`, `Description`, `AmountPerAccount`, `MaxAccounts`, `AdminEmail`, `SourceKey`) VALUES ('SKY test promo', 'SKY test promo', '1', '20', 'sky.test@test.com', '123456');

INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '1');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '2');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '3');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '4');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '5');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '6');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '7');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '8');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '9');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '10');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '11');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '12');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '13');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '14');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '15');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '16');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '17');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '18');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '19');
INSERT INTO `PromoCode` (`PromoId`, `Code`) VALUES ('1', '20');

INSERT INTO `Registration` (`PromoId`, `PromoCodeId`, `FirstName`, `LastName`, `Email`, `Mobile`, `AddressLine1`, `AddressLine2`, `City`, `State`, `Postcode`, `IP`, `UserAgent`, `CountryCode`) VALUES ('1', '1', 'Test', 'User', 'test@user.com', '123', 'Test', 'Test', 'Test', 'Test', 'Test', '111.111.111.111', 'TestUserAgent', 'US');
