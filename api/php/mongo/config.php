<?php

/**
 * by default connects to localhost
 */

$connection = new MongoClient();

/**
 * Selects the db
 */

$db = $connection -> mydb;
?>
