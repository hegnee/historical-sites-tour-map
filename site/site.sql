-- site/site.sql

-- Sites Table: Stores core information about historical sites
CREATE TABLE `sites` (
    `id` VARCHAR(255) NOT NULL PRIMARY KEY COMMENT 'Unique identifier for the site (e.g., UUID)',
    `name` VARCHAR(255) NOT NULL COMMENT 'Name of the site',
    `description` TEXT COMMENT 'Detailed description of the site',
    `category` VARCHAR(100) COMMENT 'Category of the site (e.g., Museum, Monument, Landmark)',
    `region` VARCHAR(100) COMMENT 'Geographical area or city where the site is located',
    `historical_period` VARCHAR(100) COMMENT 'Historical period the site belongs to (e.g., Ming Dynasty, Roman Empire)',
    `latitude` DECIMAL(10, 8) COMMENT 'Latitude coordinate of the site',
    `longitude` DECIMAL(11, 8) COMMENT 'Longitude coordinate of the site',
    `address` VARCHAR(255) COMMENT 'Full address of the site',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of when the record was created',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp of when the record was last updated',
    INDEX `idx_category` (`category`),
    INDEX `idx_region` (`region`),
    INDEX `idx_historical_period` (`historical_period`),
    FULLTEXT INDEX `ft_name_description` (`name`, `description`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Stores information about historical sites';

-- Multimedia Table: Stores links and information about multimedia resources related to sites
CREATE TABLE `multimedia` (
    `id` VARCHAR(255) NOT NULL PRIMARY KEY COMMENT 'Unique identifier for the multimedia resource (e.g., UUID)',
    `site_id` VARCHAR(255) NOT NULL COMMENT 'Foreign key referencing the sites table',
    `type` VARCHAR(50) NOT NULL COMMENT 'Type of multimedia (e.g., image, video, audio)',
    `url` VARCHAR(2048) NOT NULL COMMENT 'URL or path to the multimedia file (managed by media service)',
    `description` VARCHAR(500) COMMENT 'Description of the multimedia resource',
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Timestamp of when the record was created',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp of when the record was last updated',
    FOREIGN KEY (`site_id`) REFERENCES `sites`(`id`) ON DELETE CASCADE,
    INDEX `idx_site_id_type` (`site_id`, `type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Stores multimedia resources related to sites';

-- Dynamic Info Table: Stores information that changes frequently, like opening hours and ticket prices
CREATE TABLE `dynamic_info` (
    `id` VARCHAR(255) NOT NULL PRIMARY KEY COMMENT 'Unique identifier for the dynamic info record (e.g., UUID)',
    `site_id` VARCHAR(255) NOT NULL UNIQUE COMMENT 'Foreign key referencing the sites table, unique as one dynamic info per site',
    `opening_hours` JSON COMMENT 'Structured opening hours (e.g., {"weekday": "9am-5pm", "weekend": "10am-4pm"})',
    `ticket_price` VARCHAR(255) COMMENT 'Ticket price information (e.g., "Adult: $20, Child: $10")',
    `notices` JSON COMMENT 'Array of important notices or announcements (e.g., ["Closed for renovation", "Special event this weekend"])',
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Timestamp of when this information was last updated',
    FOREIGN KEY (`site_id`) REFERENCES `sites`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='Stores dynamic information for sites like opening hours and ticket prices';
