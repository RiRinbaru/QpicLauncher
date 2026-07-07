<script setup lang="ts">
import { ref } from "vue";
// 分割した部品（コンポーネント）を読み込む
import LibraryView from "./components/LibraryView.vue";
import SettingsView from "./components/SettingsView.vue";

const currentTab = ref<"library" | "settings">("library");
</script>

<template>
  <div class="launcher-container">
    <header class="top-navigation">
      <button :class="{ active: currentTab === 'library' }" @click="currentTab = 'library'">
        ライブラリ
      </button>
      <button :class="{ active: currentTab === 'settings' }" @click="currentTab = 'settings'">
        設定
      </button>
    </header>

    <main class="content-area">
      <LibraryView v-if="currentTab === 'library'" />
      <SettingsView v-if="currentTab === 'settings'" />
    </main>
  </div>
</template>

<style scoped>
/* 画面全体の大枠のスタイルのみを残す */
.launcher-container {
  display: flex;
  flex-direction: column;
  height: 100vh;
  font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
  background-color: #1e1e1e;
  color: #ffffff;
}

.top-navigation {
  display: flex;
  background-color: #2d2d2d;
  padding: 10px 20px;
  border-bottom: 1px solid #3d3d3d;
  gap: 10px;
}

.top-navigation button {
  background: none;
  border: none;
  color: #aaaaaa;
  padding: 8px 16px;
  font-size: 16px;
  cursor: pointer;
  border-radius: 4px;
  transition: all 0.2s;
}

.top-navigation button:hover {
  color: #ffffff;
  background-color: #3d3d3d;
}

.top-navigation button.active {
  color: #ffffff;
  background-color: #4a4a4a;
  font-weight: bold;
}

.content-area {
  flex: 1;
  padding: 20px;
  overflow-y: auto;
}
</style>

<style>
body {
  margin: 0;
  padding: 0;
  box-sizing: border-box;
  background-color: #1e1e1e; /* スクロール時の背景色も統一 */
}
</style>