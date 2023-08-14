CREATE USER interview_apps_user WITH CREATEDB NOSUPERUSER INHERIT PASSWORD 'P@ssw0rd';
CREATE DATABASE interview_apps_db OWNER interview_apps_user;

\c interview_apps_db interview_apps_user

CREATE TABLE university (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE major (
  id VARCHAR(100) PRIMARY KEY,
  name varchar(100) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TYPE user_role AS ENUM ('admin', 'recruiter', 'interviewer');

CREATE TABLE users (
    id VARCHAR(100) PRIMARY KEY,
    username VARCHAR(100) NOT NULL,
    password VARCHAR(100) NOT NULL,
    role user_role,
    is_active bool,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE candidate (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE candidate_resume (
    id VARCHAR(100) PRIMARY KEY,
    candidate_id VARCHAR(100),
    university_id VARCHAR(100),
    major_id VARCHAR(100),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (candidate_id) REFERENCES candidate(id),
    FOREIGN KEY (university_id) REFERENCES university(id),
    FOREIGN KEY (major_id) REFERENCES major(id)
);

CREATE TABLE interviewer (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  user_id VARCHAR(100),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE recruiter (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  user_id VARCHAR(100),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE bootcamp_source (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TYPE bootcamp_process_status_type AS ENUM ('schedule', 'interview');

CREATE TABLE bootcamp_process_status (
  id VARCHAR(100) PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  status bootcamp_process_status_type NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE interview_book (
  id VARCHAR(100) PRIMARY KEY,
  interview_date DATE,
  interviewer_id VARCHAR(100),
  recruiter_id VARCHAR(100),
  bootcamp_resource_id VARCHAR(100),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (interviewer_id) REFERENCES interviewer(id),
  FOREIGN KEY (recruiter_id) REFERENCES recruiter(id),
  FOREIGN KEY (bootcamp_resource_id) REFERENCES bootcamp_source(id)
);

CREATE TABLE interview_book_detail (
  id VARCHAR(100) PRIMARY KEY,
  interview_book_id VARCHAR(100),
  interview_time TIME,
  candidate_resume_id VARCHAR(100),
  interview_file VARCHAR(255),
  meeting_link VARCHAR(255),
  interview_status_id VARCHAR(100),
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (interview_book_id) REFERENCES interview_book(id),
  FOREIGN KEY (candidate_resume_id) REFERENCES candidate_resume(id),
  FOREIGN KEY (interview_status_id) REFERENCES bootcamp_process_status(id)
);

CREATE TABLE interview_form (
  id VARCHAR(100) PRIMARY KEY,
  title VARCHAR(100) not null,
  description TEXT,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP
);

CREATE TABLE interview_form_point (
    id VARCHAR(100) PRIMARY KEY,
    interview_form_id VARCHAR(100),
    point INT,
    title VARCHAR(255),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (interview_form_id) REFERENCES interview_form(id)
);

CREATE TABLE interview_form_result (
  id VARCHAR(100) PRIMARY KEY,
  interview_form_id VARCHAR(100),
  interview_form_point_id VARCHAR(100),
  note TEXT,
  candidate_resume_id VARCHAR(100),
  interviewer_id VARCHAR(100),
  result VARCHAR(100),
  interviewer_note text,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP,
  FOREIGN KEY (interview_form_id) REFERENCES interview_form(id),
  FOREIGN KEY (interview_form_point_id) REFERENCES interview_form_point(id),
  FOREIGN KEY (candidate_resume_id) REFERENCES candidate_resume(id),
  FOREIGN KEY (interviewer_id) REFERENCES interviewer(id),
  FOREIGN KEY (result) REFERENCES bootcamp_process_status(id)
);