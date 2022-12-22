<?php

include_once __DIR__ . '/config.php';

class DatabaseConnection
{
    private $host = DB_HOST;
    private $username = DB_USERNAME;
    private $password = DB_PW;
    private $database = DB_DATABASE;

    private $conn;

    public function __construct()
    {
        // Verbindung herstellen
        $this->conn = new mysqli($this->host, $this->username, $this->password, $this->database);

        // Überprüfen, ob die Verbindung erfolgreich hergestellt wurde
        if ($this->conn->connect_error) {
            die("Verbindung fehlgeschlagen: " . $this->conn->connect_error);
        }
    }

    public function getConnection()
    {
        return $this->conn;
    }
}

$database = new DatabaseConnection();
$conn = $database->getConnection();
echo "Verbindung erfolgreich hergestellt";
?>
