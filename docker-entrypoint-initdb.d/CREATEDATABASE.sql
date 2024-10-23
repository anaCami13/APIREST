CREATE DATABASE IF NOT EXISTS API_REST;

USE API_REST;

CREATE TABLE `actividades` (
    `id_act` INT(11) NOT NULL AUTO_INCREMENT,
    `nombre` VARCHAR(50) NULL,
    `contenido` TEXT NULL,
    PRIMARY KEY (`id_act`)
) ENGINE = InnoDB;

INSERT INTO `actividades` (`id_act`, `nombre`, `contenido`) VALUES 
(1, 'Actividad Uno', 'Caminata bajo la luna y las estrellas'), 
(2, 'Actividad Dos', 'Ver familia de diez en la TV con Arturo'), 
(3, 'Actividad Tres', 'Escribir artículos de alta calidad');

