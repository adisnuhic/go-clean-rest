INSERT INTO users (id, email, is_confirmed, accepted_tos) 
VALUES (1, "admin@admin.com", true, true);


-- pass: root
INSERT INTO auth_providers (provider, user_id, uid) VALUES
    ("local",1,"$2a$12$c1LYye8FDfcTugQs05VqhelEcL30up4N/opyelt.bQcpK/v8P9aeu");
