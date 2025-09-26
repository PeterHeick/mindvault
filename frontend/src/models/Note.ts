// frontend/src/models/Note.ts
export interface Note {
  id: number;
  title: string;
  content: string;
  created_at: string; // Datoer kommer som strenge fra JSON
  updated_at: string;
}