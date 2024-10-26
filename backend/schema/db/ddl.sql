CREATE TABLE "echos"
(
    "id"         UUID         NOT NULL,
    "message"    VARCHAR(255) NOT NULL,
    "timestamp"  TIMESTAMP    NOT NULL,
    "created_at" TIMESTAMP    NOT NULL,
    "updated_at" TIMESTAMP    NOT NULL,
    PRIMARY KEY ("id")
);

-- music_sheets テーブルを作成
CREATE TABLE "music_sheets"
(
    "music_sheet_id" UUID PRIMARY KEY,
    "title"          VARCHAR(100) NOT NULL,
    "created_at"     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- notes テーブルを作成
CREATE TABLE "notes"
(
    "note_id"        UUID PRIMARY KEY,
    "music_sheet_id" UUID NOT NULL REFERENCES "music_sheets"("music_sheet_id") ON DELETE CASCADE,
    "pitches"        INT[] NOT NULL,
    "created_at"     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"     TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
