-- Create a new database called 'sockets-server'
-- Connect to the 'master' database to run this snippet
-- Create the new database if it does not exist already

CREATE DATABASE
IF
	NOT EXISTS `socket-server` DEFAULT CHARACTER 
	SET utf8mb4 COLLATE utf8mb4_unicode_ci;