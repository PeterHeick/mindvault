<script setup lang="ts">
import { ref, onMounted } from 'vue';
import NoteService from '@/services/NoteService';
import type { Note } from '@/models/Note';

const notes = ref<Note[]>([]);
const error = ref<string | null>(null);

// Refs til vores nye note-formular
const newNoteTitle = ref('');
const newNoteContent = ref('');

// Hent alle noter, når komponenten indlæses
onMounted(async () => {
  try {
    const response = await NoteService.getNotes();
    notes.value = response.data;
  } catch (err) {
    error.value = 'Could not fetch notes from the API.';
    console.error(err);
  }
});

async function createNote() {
  if (!newNoteTitle.value) {
    alert('Titel kan ikke være tom.');
    return;
  }
  try {
    const response = await NoteService.createNote({
      title: newNoteTitle.value,
      content: newNoteContent.value
    });
    // Tilføj den nye note øverst i listen for øjeblikkelig feedback
    notes.value.unshift(response.data); 
    
    // Nulstil formularfelterne
    newNoteTitle.value = '';
    newNoteContent.value = '';
  } catch (err) {
    alert('Fejl ved oprettelse af note.');
    console.error(err);
  }
}

async function deleteNote(id: number) {
  if (!confirm('Er du sikker på, du vil slette denne note?')) {
    return;
  }
  try {
    await NoteService.deleteNote(id);
    // Fjern noten fra listen i UI for øjeblikkelig feedback
    notes.value = notes.value.filter(note => note.id !== id);
  } catch (err) {
    alert('Fejl ved sletning af note.');
    console.error(err);
  }
}
</script>

<template>
  <div class="notes-app">
    <h1>MindVault Noter</h1>
    
    <form @submit.prevent="createNote" class="note-form">
      <input type="text" v-model="newNoteTitle" placeholder="Titel på ny note..." />
      <textarea v-model="newNoteContent" placeholder="Skriv din note her..."></textarea>
      <button type="submit">Opret Note</button>
    </form>

    <div v-if="error" class="error">{{ error }}</div>
    
    <ul v-if="notes.length" class="notes-list">
      <li v-for="note in notes" :key="note.id">
        <h2>{{ note.title }}</h2>
        <p>{{ note.content }}</p>
        <div class="note-footer">
          <small>Opdateret: {{ new Date(note.updated_at).toLocaleString() }}</small>
          <button @click="deleteNote(note.id)" class="delete-btn">Slet</button>
        </div>
      </li>
    </ul>
    <p v-else-if="!error">Du har ingen noter endnu. Opret din første!</p>
  </div>
</template>

<style scoped>
.notes-app { max-width: 700px; margin: 0 auto; }
.note-form {
  display: flex;
  flex-direction: column;
  margin-bottom: 2rem;
  gap: 0.5rem;
}
.note-form input, .note-form textarea {
  padding: 0.75rem;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 1rem;
}
.note-form button {
  padding: 0.75rem;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  cursor: pointer;
  font-size: 1rem;
}
.note-form button:hover { background-color: #0056b3; }

.notes-list { list-style: none; padding: 0; }
.notes-list li { background-color: #f9f9f9; padding: 1.5rem; margin-bottom: 1rem; border-radius: 5px; border: 1px solid #eee; }
.note-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-top: 1rem;
  color: #666;
  font-size: 0.8rem;
}
.delete-btn {
  background: none;
  border: 1px solid #ff4d4d;
  color: #ff4d4d;
  padding: 0.3rem 0.6rem;
  border-radius: 5px;
  cursor: pointer;
}
.delete-btn:hover { background: #ff4d4d; color: white; }
.error { color: red; }
</style>