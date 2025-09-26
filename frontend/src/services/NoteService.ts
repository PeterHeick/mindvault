// frontend/src/services/NoteService.ts
import axios from 'axios';
import type { Note } from '@/models/Note';

const apiClient = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  headers: {
    'Content-Type': 'application/json'
  }
});

export default {
  getNotes() {
    return apiClient.get<Note[]>('/notes');
  },
  createNote(note: { title: string; content: string }) {
    return apiClient.post<Note>('/notes', note);
  },
  // Vi tilføjer 'update' og 'delete' senere for at holde det simpelt
  deleteNote(id: number) {
    return apiClient.delete(`/notes/${id}`);
  }
};