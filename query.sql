/* =========================================================
   QUERY PUBLIC – BIODATA
   Mengambil data biodata utama pengguna
   Digunakan di Hero & About Section
   ========================================================= */
SELECT
    id,
    user_id,
    full_name,
    title,
    about,
    address,
    profile_image,
    cv_url
FROM biodata
WHERE user_id = $1;


/* =========================================================
   QUERY PUBLIC – SOCIAL MEDIA
   Mengambil daftar media sosial pengguna
   Digunakan di Hero & Footer
   ========================================================= */
SELECT
    id,
    platform,
    url,
    icon
FROM social_media
WHERE user_id = $1
ORDER BY id;


/* =========================================================
   QUERY PUBLIC – EDUCATION
   Mengambil riwayat pendidikan (timeline)
   ========================================================= */
SELECT
    id,
    institution,
    degree,
    major,
    start_year,
    end_year,
    description
FROM education
WHERE user_id = $1
ORDER BY start_year DESC;


/* =========================================================
   QUERY PUBLIC – SKILLS
   Mengambil daftar skill pengguna
   ========================================================= */
SELECT
    id,
    name,
    category,
    level,
    icon
FROM skills
WHERE user_id = $1
ORDER BY category, name;


/* =========================================================
   QUERY PUBLIC – EXPERIENCE
   Mengambil pengalaman kerja / organisasi
   ========================================================= */
SELECT
    id,
    title,
    company,
    type,
    start_date,
    end_date,
    is_current,
    description
FROM experience
WHERE user_id = $1
ORDER BY start_date DESC;


/* =========================================================
   QUERY PUBLIC – FEATURED PROJECTS
   Mengambil project unggulan (homepage)
   ========================================================= */
SELECT
    id,
    title,
    slug,
    description,
    thumbnail,
    tech_stack,
    github_url,
    demo_url
FROM projects
WHERE user_id = $1
  AND is_featured = TRUE
ORDER BY created_at DESC;


/* =========================================================
   QUERY PUBLIC – ALL PROJECTS
   Mengambil semua project untuk halaman portfolio
   ========================================================= */
SELECT
    id,
    title,
    slug,
    description,
    thumbnail,
    tech_stack,
    github_url,
    demo_url
FROM projects
WHERE user_id = $1
ORDER BY created_at DESC;


/* =========================================================
   QUERY PUBLIC – PROJECT DETAIL
   Mengambil detail satu project berdasarkan slug
   ========================================================= */
SELECT
    id,
    title,
    description,
    thumbnail,
    tech_stack,
    github_url,
    demo_url,
    created_at
FROM projects
WHERE slug = $1;


/* =========================================================
   QUERY PUBLIC – INSERT CONTACT
   Menyimpan pesan dari form contact
   ========================================================= */
INSERT INTO contacts (
    name,
    email,
    subject,
    message
) VALUES (
    $1, $2, $3, $4
);


/* =========================================================
   QUERY ADMIN – LOGIN
   Digunakan untuk autentikasi admin
   ========================================================= */
SELECT
    id,
    username,
    password_hash,
    role,
    is_active
FROM users
WHERE username = $1;


/* =========================================================
   QUERY ADMIN – DASHBOARD STATISTIC
   Menampilkan ringkasan data di dashboard admin
   ========================================================= */
SELECT
    (SELECT COUNT(*) FROM projects WHERE user_id = $1)     AS total_projects,
    (SELECT COUNT(*) FROM skills WHERE user_id = $1)       AS total_skills,
    (SELECT COUNT(*) FROM experience WHERE user_id = $1)   AS total_experience,
    (SELECT COUNT(*) FROM contacts)                         AS total_messages;


/* =========================================================
   QUERY ADMIN – LIST PROJECT
   Menampilkan daftar project di dashboard admin
   ========================================================= */
SELECT
    id,
    title,
    slug,
    is_featured,
    created_at
FROM projects
WHERE user_id = $1
ORDER BY created_at DESC;


/* =========================================================
   QUERY ADMIN – CREATE PROJECT
   Menambahkan project baru
   ========================================================= */
INSERT INTO projects (
    user_id,
    title,
    slug,
    description,
    thumbnail,
    tech_stack,
    github_url,
    demo_url,
    is_featured
) VALUES (
    $1, $2, $3, $4,
    $5, $6, $7, $8, $9
);


/* =========================================================
   QUERY ADMIN – GET PROJECT BY ID
   Mengambil data project untuk proses edit
   ========================================================= */
SELECT
    id,
    title,
    slug,
    description,
    thumbnail,
    tech_stack,
    github_url,
    demo_url,
    is_featured
FROM projects
WHERE id = $1 AND user_id = $2;


/* =========================================================
   QUERY ADMIN – UPDATE PROJECT
   Memperbarui data project
   ========================================================= */
UPDATE projects SET
    title = $1,
    slug = $2,
    description = $3,
    thumbnail = $4,
    tech_stack = $5,
    github_url = $6,
    demo_url = $7,
    is_featured = $8
WHERE id = $9 AND user_id = $10;


/* =========================================================
   QUERY ADMIN – DELETE PROJECT
   Menghapus project
   ========================================================= */
DELETE FROM projects
WHERE id = $1 AND user_id = $2;


/* =========================================================
   QUERY ADMIN – SOCIAL MEDIA LIST
   ========================================================= */
SELECT
    id,
    platform,
    url,
    icon
FROM social_media
WHERE user_id = $1;


/* =========================================================
   QUERY ADMIN – CREATE SOCIAL MEDIA
   ========================================================= */
INSERT INTO social_media (
    user_id,
    platform,
    url,
    icon
) VALUES (
    $1, $2, $3, $4
);


/* =========================================================
   QUERY ADMIN – UPDATE SOCIAL MEDIA
   ========================================================= */
UPDATE social_media SET
    platform = $1,
    url = $2,
    icon = $3
WHERE id = $4 AND user_id = $5;


/* =========================================================
   QUERY ADMIN – DELETE SOCIAL MEDIA
   ========================================================= */
DELETE FROM social_media
WHERE id = $1 AND user_id = $2;
