-- This script only contains the table creation statements and does not fully represent the table in the database. Do not use it as a backup.

-- Sequence and defined type
CREATE SEQUENCE IF NOT EXISTS roles_id_seq;

-- Table Definition
CREATE TABLE "public"."roles" (
    "id" int4 NOT NULL DEFAULT nextval('roles_id_seq'::regclass),
    "name" varchar,
    "language" varchar,
    PRIMARY KEY ("id")
);

INSERT INTO
    roles (name, language)
VALUES ('วิศวกร', 'th'), -- Engineer
    ('หมอ', 'th'), -- Doctor
    ('ทนายความ', 'th'), -- Lawyer
    ('นักวิทยาศาสตร์', 'th'), -- Scientist
    ('ครู', 'th'), -- Teacher
    ('นักเรียน', 'th'), -- Student
    ('นักธุรกิจ', 'th'), -- Businessman
    ('นักการเมือง', 'th'), -- Politician
    ('ชาวนา', 'th'), -- Farmer
    ('ศิลปิน', 'th'), -- Artist
    ('นักดนตรี', 'th'), -- Musician
    ('นักแสดง', 'th'), -- Actor
    ('นักกีฬา', 'th'), -- Athlete
    ('ทหาร', 'th'), -- Soldier
    ('ตำรวจ', 'th'), -- Police
    ('นักดับเพลิง', 'th'), -- Firefighter
    ('เชฟ', 'th'), -- Chef
    ('นักข่าว', 'th'), -- Journalist
    ('นักเขียน', 'th'), -- Writer
    ('ช่างภาพ', 'th'), -- Photographer
    ('นักออกแบบ', 'th'), -- Designer
    ('สถาปนิก', 'th'), -- Architect
    ('นักเต้น', 'th'), -- Dancer
    ('นักร้อง', 'th'), -- Singer
    ('นางแบบ', 'th'), -- Model
    ('ผู้กำกับ', 'th'), -- Director
    ('โปรดิวเซอร์', 'th'), -- Producer
    ('ช่างกล้อง', 'th'), -- Cameraman
    ('บรรณาธิการ', 'th'), -- Editor
    ('ช่างภาพยนตร์', 'th'), -- Cinematographer
    ('วิศวกรเสียง', 'th'), -- Sound Engineer
    ('ศิลปิน VFX', 'th'), -- VFX Artist
    ('นักอนิเมชั่น', 'th'), -- Animator
    ('นักพัฒนาเกม', 'th'), -- Game Developer
    ('นักพัฒนาเว็บ', 'th'), -- Web Developer
    ('นักพัฒนาแอป', 'th'), -- App Developer
    ('นักพัฒนาซอฟต์แวร์', 'th'), -- Software Developer
    ('วิศวกรฮาร์ดแวร์', 'th'), -- Hardware Engineer
    ('วิศวกรเครือข่าย', 'th'), -- Network Engineer
    ('ผู้ดูแลฐานข้อมูล', 'th'), -- Database Administrator
    ('ผู้ดูแลระบบ', 'th'), -- System Administrator
    (
        'ผู้เชี่ยวชาญด้านความปลอดภัยไซเบอร์',
        'th'
    ), -- Cyber Security Expert
    ('นักวิทยาศาสตร์ข้อมูล', 'th'), -- Data Scientist
    (
        'วิศวกรการเรียนรู้ของเครื่อง',
        'th'
    ), -- Machine Learning Engineer
    (
        'ผู้เชี่ยวชาญด้านปัญญาประดิษฐ์',
        'th'
    ), -- Artificial Intelligence Expert
    ('วิศวกรหุ่นยนต์', 'th'), -- Robotics Engineer
    ('วิศวกรระบบอัตโนมัติ', 'th'), -- Automation Engineer
    ('วิศวกรไฟฟ้า', 'th'), -- Electrical Engineer
    ('วิศวกรเครื่องกล', 'th'), -- Mechanical Engineer
    ('วิศวกรโยธา', 'th'), -- Civil Engineer
    ('วิศวกรเคมี', 'th'), -- Chemical Engineer
    ('วิศวกรอากาศยาน', 'th'), -- Aerospace Engineer
    ('วิศวกรยานยนต์', 'th'), -- Automobile Engineer
    ('วิศวกรทางทะเล', 'th'), -- Marine Engineer
    ('วิศวกรเหมืองแร่', 'th'), -- Mining Engineer
    ('วิศวกรปิโตรเลียม', 'th'), -- Petroleum Engineer
    ('วิศวกรสิ่งแวดล้อม', 'th'), -- Environmental Engineer
    ('วิศวกรชีวการแพทย์', 'th'), -- Biomedical Engineer
    ('นักชีวเทคโนโลยี', 'th'), -- Biotechnologist
    ('นักพันธุศาสตร์', 'th'), -- Geneticist
    ('นักจุลชีววิทยา', 'th'), -- Microbiologist
    ('นักชีวเคมี', 'th'), -- Biochemist
    ('เภสัชกร', 'th'), -- Pharmacist
    ('นักกายภาพบำบัด', 'th'), -- Physiotherapist
    ('พยาบาล', 'th'), -- Nurse
    ('ทันตแพทย์', 'th'), -- Dentist
    ('สัตวแพทย์', 'th'), -- Veterinarian
    ('นักจิตวิทยา', 'th'), -- Psychologist
    ('จิตแพทย์', 'th'), -- Psychiatrist
    ('ที่ปรึกษา', 'th'), -- Counselor
    ('นักสังคมสงเคราะห์', 'th'), -- Social Worker
    (
        'ผู้จัดการฝ่ายทรัพยากรบุคคล',
        'th'
    ), -- Human Resource Manager
    ('ผู้จัดการฝ่ายการตลาด', 'th'), -- Marketing Manager
    ('ผู้จัดการฝ่ายขาย', 'th'), -- Sales Manager
    ('ผู้จัดการฝ่ายการเงิน', 'th'), -- Finance Manager
    ('นักบัญชี', 'th'), -- Accountant
    ('ผู้ตรวจสอบบัญชี', 'th'), -- Auditor
    ('ที่ปรึกษาด้านภาษี', 'th'), -- Tax Consultant
    ('นักวิเคราะห์การลงทุน', 'th'), -- Investment Banker
    (
        'นายหน้าซื้อขายหลักทรัพย์',
        'th'
    ), -- Stock Broker
    ('ตัวแทนประกันภัย', 'th'), -- Insurance Agent
    (
        'นายหน้าอสังหาริมทรัพย์',
        'th'
    ), -- Real Estate Agent
    ('ตัวแทนท่องเที่ยว', 'th'), -- Travel Agent
    ('ผู้จัดการงานอีเว้นท์', 'th'), -- Event Manager
    ('ไกด์นำเที่ยว', 'th'), -- Tour Guide
    ('ผู้จัดการโรงแรม', 'th'), -- Hotel Manager
    ('ผู้จัดการร้านอาหาร', 'th'), -- Restaurant Manager
    ('คนทำขนมปัง', 'th'), -- Baker
    ('บาริสต้า', 'th'), -- Barista
    ('บาร์เทนเดอร์', 'th'), -- Bartender
    ('พนักงานเสิร์ฟ', 'th'), -- Waiter
    ('พนักงานทำความสะอาด', 'th'), -- Housekeeper
    ('ยาม', 'th'), -- Security Guard
    ('คนสวน', 'th'), -- Gardener
    ('คนขับรถ', 'th'), -- Driver
    ('ช่างเครื่องยนต์', 'th'), -- Mechanic
    ('ช่างประปา', 'th'), -- Plumber
    ('ช่างไฟฟ้า', 'th'), -- Electrician
    ('ช่างไม้', 'th'), -- Carpenter
    ('ช่างก่ออิฐ', 'th'), -- Mason
    ('ช่างเชื่อม', 'th'), -- Welder
    ('ช่างเหล็ก', 'th'), -- Blacksmith
    ('ช่างตัดเสื้อ', 'th'), -- Tailor
    ('ช่างทำรองเท้า', 'th'), -- Shoemaker
    ('ช่างทำเครื่องประดับ', 'th'), -- Jeweler
    ('ช่างทำนาฬิกา', 'th'), -- Watchmaker
    ('จิตรกร', 'th'), -- Painter
    ('ช่างแกะสลัก', 'th'), -- Sculptor
    ('ช่างปั้นหม้อ', 'th'), -- Potter
    ('ช่างทอผ้า', 'th'), -- Weaver
    ('ช่างสานตะกร้า', 'th'), -- Basket Weaver
    ('นักบาสเกตบอล', 'th'), -- Basketball Player
    ('นักฟุตบอล', 'th'), -- Football Player
    ('นักคริกเก็ต', 'th'), -- Cricket Player
    ('นักเทนนิส', 'th'), -- Tennis Player
    ('นักกอล์ฟ', 'th'), -- Golf Player
    ('นักว่ายน้ำ', 'th'), -- Swimmer
    ('นักวิ่ง', 'th'), -- Runner
    ('นักปั่นจักรยาน', 'th'), -- Cyclist
    ('นักมวย', 'th'), -- Boxer
    ('นักมวยปล้ำ', 'th'), -- Wrestler
    ('นักยกน้ำหนัก', 'th'), -- Weightlifter
    ('นักยิมนาสติก', 'th'), -- Gymnast
    ('ครูสอนโยคะ', 'th'), -- Yoga Instructor
    ('นักศิลปะการต่อสู้', 'th'), -- Martial Artist
    ('โค้ช', 'th'), -- Coach
    ('ผู้ตัดสิน', 'th'), -- Referee
    ('กรรมการ', 'th'), -- Umpire
    ('นักข่าวกีฬา', 'th'), -- Sports Journalist
    ('ผู้บรรยายกีฬา', 'th'), -- Sports Commentator
    ('นักวิเคราะห์กีฬา', 'th'), -- Sports Analyst
    ('ผู้ถ่ายทอดกีฬา', 'th'), -- Sports Broadcaster
    ('ช่างภาพกีฬา', 'th'), -- Sports Photographer
    ('ช่างวิดีโอกีฬา', 'th'), -- Sports Videographer
    ('บรรณาธิการกีฬา', 'th'), -- Sports Editor
    ('โปรดิวเซอร์กีฬา', 'th'), -- Sports Producer
    ('ผู้กำกับกีฬา', 'th'), -- Sports Director
    ('ช่างกล้องกีฬา', 'th'), -- Sports Cameraman
    ('วิศวกรเสียงกีฬา', 'th'), -- Sports Sound Engineer
    ('ศิลปิน VFX กีฬา', 'th'), -- Sports VFX Artist
    ('นักอนิเมชั่นกีฬา', 'th'), -- Sports Animator
    ('นักพัฒนาเกมกีฬา', 'th'), -- Sports Game Developer
    ('นักพัฒนาเว็บกีฬา', 'th'), -- Sports Web Developer
    ('นักพัฒนาแอปกีฬา', 'th'), -- Sports App Developer
    ('นักพัฒนาซอฟต์แวร์กีฬา', 'th'), -- Sports Software Developer
    ('วิศวกรฮาร์ดแวร์กีฬา', 'th'), -- Sports Hardware Engineer
    ('วิศวกรเครือข่ายกีฬา', 'th'), -- Sports Network Engineer
    ('ผู้ดูแลฐานข้อมูลกีฬา', 'th'), -- Sports Database Administrator
    ('ผู้ดูแลระบบกีฬา', 'th'), -- Sports System Administrator
    (
        'ผู้เชี่ยวชาญด้านความปลอดภัยไซเบอร์กีฬา',
        'th'
    ), -- Sports Cyber Security Expert
    (
        'นักวิทยาศาสตร์ข้อมูลกีฬา',
        'th'
    ), -- Sports Data Scientist
    (
        'วิศวกรการเรียนรู้ของเครื่องกีฬา',
        'th'
    ), -- Sports Machine Learning Engineer
    (
        'ผู้เชี่ยวชาญด้านปัญญาประดิษฐ์กีฬา',
        'th'
    ), -- Sports Artificial Intelligence Expert
    ('วิศวกรหุ่นยนต์กีฬา', 'th'), -- Sports Robotics Engineer
    (
        'วิศวกรระบบอัตโนมัติกีฬา',
        'th'
    ), -- Sports Automation Engineer
    ('วิศวกรไฟฟ้ากีฬา', 'th'), -- Sports Electrical Engineer
    ('วิศวกรเครื่องกลกีฬา', 'th'), -- Sports Mechanical Engineer
    ('วิศวกรโยธากีฬา', 'th'), -- Sports Civil Engineer
    ('วิศวกรเคมีกีฬา', 'th'), -- Sports Chemical Engineer
    ('วิศวกรอากาศยานกีฬา', 'th'), -- Sports Aerospace Engineer
    ('วิศวกรยานยนต์กีฬา', 'th');
-- Sports Automobile Engineer


-- Insert roles with default language set to 'en'
INSERT INTO
    roles (name, language)
VALUES ('Engineer', 'en'),
    ('Doctor', 'en'),
    ('Lawyer', 'en'),
    ('Scientist', 'en'),
    ('Teacher', 'en'),
    ('Student', 'en'),
    ('Businessman', 'en'),
    ('Politician', 'en'),
    ('Farmer', 'en'),
    ('Artist', 'en'),
    ('Musician', 'en'),
    ('Actor', 'en'),
    ('Athlete', 'en'),
    ('Soldier', 'en'),
    ('Police', 'en'),
    ('Firefighter', 'en'),
    ('Chef', 'en'),
    ('Journalist', 'en'),
    ('Writer', 'en'),
    ('Photographer', 'en'),
    ('Designer', 'en'),
    ('Architect', 'en'),
    ('Dancer', 'en'),
    ('Singer', 'en'),
    ('Model', 'en'),
    ('Director', 'en'),
    ('Producer', 'en'),
    ('Cameraman', 'en'),
    ('Editor', 'en'),
    ('Cinematographer', 'en'),
    ('Sound Engineer', 'en'),
    ('VFX Artist', 'en'),
    ('Animator', 'en'),
    ('Game Developer', 'en'),
    ('Web Developer', 'en'),
    ('App Developer', 'en'),
    ('Software Developer', 'en'),
    ('Hardware Engineer', 'en'),
    ('Network Engineer', 'en'),
    (
        'Database Administrator',
        'en'
    ),
    ('System Administrator', 'en'),
    ('Cyber Security Expert', 'en'),
    ('Data Scientist', 'en'),
    (
        'Machine Learning Engineer',
        'en'
    ),
    (
        'Artificial Intelligence Expert',
        'en'
    ),
    ('Robotics Engineer', 'en'),
    ('Automation Engineer', 'en'),
    ('Electrical Engineer', 'en'),
    ('Mechanical Engineer', 'en'),
    ('Civil Engineer', 'en'),
    ('Chemical Engineer', 'en'),
    ('Aerospace Engineer', 'en'),
    ('Automobile Engineer', 'en'),
    ('Marine Engineer', 'en'),
    ('Mining Engineer', 'en'),
    ('Petroleum Engineer', 'en'),
    (
        'Environmental Engineer',
        'en'
    ),
    ('Biomedical Engineer', 'en'),
    ('Biotechnologist', 'en'),
    ('Geneticist', 'en'),
    ('Microbiologist', 'en'),
    ('Biochemist', 'en'),
    ('Pharmacist', 'en'),
    ('Physiotherapist', 'en'),
    ('Nurse', 'en'),
    ('Dentist', 'en'),
    ('Veterinarian', 'en'),
    ('Psychologist', 'en'),
    ('Psychiatrist', 'en'),
    ('Counselor', 'en'),
    ('Social Worker', 'en'),
    (
        'Human Resource Manager',
        'en'
    ),
    ('Marketing Manager', 'en'),
    ('Sales Manager', 'en'),
    ('Finance Manager', 'en'),
    ('Accountant', 'en'),
    ('Auditor', 'en'),
    ('Tax Consultant', 'en'),
    ('Investment Banker', 'en'),
    ('Stock Broker', 'en'),
    ('Insurance Agent', 'en'),
    ('Real Estate Agent', 'en'),
    ('Travel Agent', 'en'),
    ('Event Manager', 'en'),
    ('Tour Guide', 'en'),
    ('Hotel Manager', 'en'),
    ('Restaurant Manager', 'en'),
    ('Baker', 'en'),
    ('Barista', 'en'),
    ('Bartender', 'en'),
    ('Waiter', 'en'),
    ('Housekeeper', 'en'),
    ('Security Guard', 'en'),
    ('Cleaner', 'en'),
    ('Gardener', 'en'),
    ('Driver', 'en'),
    ('Mechanic', 'en'),
    ('Plumber', 'en'),
    ('Electrician', 'en'),
    ('Carpenter', 'en'),
    ('Mason', 'en'),
    ('Welder', 'en'),
    ('Blacksmith', 'en'),
    ('Tailor', 'en'),
    ('Shoemaker', 'en'),
    ('Jeweler', 'en'),
    ('Watchmaker', 'en'),
    ('Painter', 'en'),
    ('Sculptor', 'en'),
    ('Potter', 'en'),
    ('Weaver', 'en'),
    ('Basket Weaver', 'en'),
    ('Basketball Player', 'en'),
    ('Football Player', 'en'),
    ('Cricket Player', 'en'),
    ('Tennis Player', 'en'),
    ('Golf Player', 'en'),
    ('Swimmer', 'en'),
    ('Runner', 'en'),
    ('Cyclist', 'en'),
    ('Boxer', 'en'),
    ('Wrestler', 'en'),
    ('Weightlifter', 'en'),
    ('Gymnast', 'en'),
    ('Yoga Instructor', 'en'),
    ('Martial Artist', 'en'),
    ('Coach', 'en'),
    ('Referee', 'en'),
    ('Umpire', 'en'),
    ('Sports Journalist', 'en'),
    ('Sports Commentator', 'en'),
    ('Sports Analyst', 'en'),
    ('Sports Broadcaster', 'en'),
    ('Sports Photographer', 'en'),
    ('Sports Videographer', 'en'),
    ('Sports Editor', 'en'),
    ('Sports Producer', 'en'),
    ('Sports Director', 'en'),
    ('Sports Cameraman', 'en'),
    ('Sports Sound Engineer', 'en'),
    ('Sports VFX Artist', 'en'),
    ('Sports Animator', 'en'),
    ('Sports Game Developer', 'en'),
    ('Sports Web Developer', 'en'),
    ('Sports App Developer', 'en'),
    (
        'Sports Software Developer',
        'en'
    ),
    (
        'Sports Hardware Engineer',
        'en'
    ),
    (
        'Sports Network Engineer',
        'en'
    ),
    (
        'Sports Database Administrator',
        'en'
    ),
    (
        'Sports System Administrator',
        'en'
    ),
    (
        'Sports Cyber Security Expert',
        'en'
    ),
    ('Sports Data Scientist', 'en'),
    (
        'Sports Machine Learning Engineer',
        'en'
    ),
    (
        'Sports Artificial Intelligence Expert',
        'en'
    ),
    (
        'Sports Robotics Engineer',
        'en'
    ),
    (
        'Sports Automation Engineer',
        'en'
    ),
    (
        'Sports Electrical Engineer',
        'en'
    ),
    (
        'Sports Mechanical Engineer',
        'en'
    ),
    ('Sports Civil Engineer', 'en'),
    (
        'Sports Chemical Engineer',
        'en'
    ),
    (
        'Sports Aerospace Engineer',
        'en'
    ),
    (
        'Sports Automobile Engineer',
        'en'
    ),
    (
        'Sports Marine Engineer',
        'en'
    ),
    (
        'Sports Mining Engineer',
        'en'
    ),
    (
        'Sports Petroleum Engineer',
        'en'
    ),
    (
        'Sports Environmental Engineer',
        'en'
    ),
    (
        'Sports Biomedical Engineer',
        'en'
    ),
    (
        'Sports Biotechnologist',
        'en'
    ),
    ('Sports Geneticist', 'en'),
    ('Sports Microbiologist', 'en'),
    ('Sports Biochemist', 'en'),
    ('Sports Pharmacist', 'en'),
    (
        'Sports Physiotherapist',
        'en'
    ),
    ('Sports Nurse', 'en'),
    ('Sports Dentist', 'en'),
    ('Sports Veterinarian', 'en'),
    ('Sports Psychologist', 'en'),
    ('Sports Psychiatrist', 'en'),
    ('Sports Counselor', 'en'),
    ('Sports Social Worker', 'en'),
    (
        'Sports Human Resource Manager',
        'en'
    );
