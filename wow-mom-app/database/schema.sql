-- MOTHERS Table
CREATE TABLE MOTHERS (
    mother_id INTEGER PRIMARY KEY AUTOINCREMENT,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    phone_number TEXT,
    address TEXT,
    city TEXT,
    state TEXT,
    zip_code TEXT,
    date_of_birth DATE,
    children_info TEXT, -- Stored as JSON or comma-separated string (e.g., "5,8" for ages)
    preferred_meeting_times TEXT, -- Stored as JSON or comma-separated string
    preferred_locations TEXT, -- Stored as JSON or comma-separated string
    account_status TEXT DEFAULT 'Active', -- e.g., Active, Inactive, Suspended
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- LEADERS Table
CREATE TABLE LEADERS (
    leader_id INTEGER PRIMARY KEY AUTOINCREMENT,
    mother_id INTEGER UNIQUE, -- Optional: If a leader is also a registered mother
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    phone_number TEXT,
    bio TEXT,
    leader_status TEXT DEFAULT 'Pending Approval', -- e.g., Active, Inactive, Pending Approval
    approved_by_admin BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (mother_id) REFERENCES MOTHERS(mother_id)
);

-- GROUPS Table
CREATE TABLE GROUPS (
    group_id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    leader_id INTEGER NOT NULL,
    max_members INTEGER DEFAULT 15,
    current_members INTEGER DEFAULT 0,
    status TEXT DEFAULT 'Active', -- e.g., Active, Inactive, Full
    meeting_details TEXT,
    meeting_frequency TEXT, -- e.g., Weekly, Bi-weekly, Monthly
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (leader_id) REFERENCES LEADERS(leader_id)
);

-- APPLICATIONS Table (Mother applying to a group)
CREATE TABLE APPLICATIONS (
    application_id INTEGER PRIMARY KEY AUTOINCREMENT,
    mother_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    application_date DATETIME DEFAULT CURRENT_TIMESTAMP,
    application_status TEXT DEFAULT 'Pending', -- e.g., Pending, Interview Scheduled, Accepted, Rejected, Active
    interview_date DATETIME,
    interview_time_slot DATETIME,
    interview_confirmed_by_mother BOOLEAN DEFAULT FALSE,
    rejection_reason TEXT,
    accepted_date DATETIME,
    activated_date DATETIME,
    notes_from_leader TEXT,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (mother_id, group_id), -- A mother can apply to a group only once
    FOREIGN KEY (mother_id) REFERENCES MOTHERS(mother_id),
    FOREIGN KEY (group_id) REFERENCES GROUPS(group_id)
);

-- NOTIFICATIONS Table
CREATE TABLE NOTIFICATIONS (
    notification_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER, -- Can be mother_id, leader_id, or admin_id
    user_type TEXT NOT NULL, -- 'Mother', 'Leader', 'Admin' to differentiate user_id context
    notification_type TEXT NOT NULL, -- e.g., 'group_application', 'interview_scheduled', 'group_message'
    message TEXT NOT NULL,
    email_sent BOOLEAN DEFAULT FALSE,
    is_read BOOLEAN DEFAULT FALSE,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- Junction table for Mothers and Groups (Many-to-Many relationship)
CREATE TABLE GROUP_MEMBERS (
    group_member_id INTEGER PRIMARY KEY AUTOINCREMENT,
    mother_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    joined_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    left_at DATETIME,
    status TEXT DEFAULT 'active', -- e.g., active, inactive, suspended
    FOREIGN KEY (mother_id) REFERENCES MOTHERS(mother_id),
    FOREIGN KEY (group_id) REFERENCES GROUPS(group_id),
    UNIQUE (mother_id, group_id) -- A mother can only be a member of a group once at a time
);