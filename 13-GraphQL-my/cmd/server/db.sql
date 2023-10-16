-- SQLite3 Script to create the database structure

-- Table: category
CREATE TABLE IF NOT EXISTS categories (
                                        id TEXT NOT NULL,
                                        name TEXT NOT NULL,
                                        description TEXT
);

-- Table: course
CREATE TABLE IF NOT EXISTS courses (
                                      id TEXT NOT NULL,
                                      name TEXT NOT NULL,
                                      description TEXT,
                                      category_id INTEGER,
                                      FOREIGN KEY (category_id) REFERENCES categories (id) ON DELETE CASCADE ON UPDATE CASCADE
    );

-- Table: lesson
CREATE TABLE IF NOT EXISTS lessons (
                                      id TEXT NOT NULL,
                                      name TEXT NOT NULL,
                                      description TEXT,
                                      content TEXT,
                                      course_id INTEGER,
                                      FOREIGN KEY (course_id) REFERENCES courses (id) ON DELETE CASCADE ON UPDATE CASCADE
    );
