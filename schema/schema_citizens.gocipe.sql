-- generated by gocipe 02fef3d117f1029d4142b6b7ae2d1ea6f313fd8a2f44e25333775a308c8afb37; DO NOT EDIT

DROP TABLE IF EXISTS citizens;

CREATE TABLE citizens (
	"id" CHAR(36),
	"surname" VARCHAR(255) NOT NULL,
	"othernames" VARCHAR(255) NOT NULL,
	"gender" CHAR(1) NOT NULL,
	"dob" DATE NOT NULL,
	"country_id" CHAR(36) NOT NULL,
	
	
	PRIMARY KEY ("id")
);

