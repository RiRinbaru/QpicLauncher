<script setup lang="ts">
import { ref } from "vue";
import { invoke } from "@tauri-apps/api/core";

// 型定義
interface Game {
  id: number;
  title: string;
  image: string;
  explanation: string;
}

//ダミーデータ
const games = ref([
  {
    id: 1,
    title: "game1",
    image: "https://picsum.photos/id/1025/300/180",
    explanation: "a",
    detail: "あいうえお",
  },
{
    id: 2,
    title: "game2",
    image: "https://picsum.photos/id/1067/300/180",
    explanation: "b",
    detail: "かきくけこ",
  },
  {
    id: 3,
    title: "game3",
    image: "https://picsum.photos/id/103/300/180",
    explanation: "c",
    detail: "さしすせそ",
  },
])

// 選択中のゲーム（nullなら一覧）
const selectedGame = ref<Game | null>(null);

// 詳細を開く
const showDetail = (game: Game) => {
  selectedGame.value = game;
};

// 詳細から戻る
const closeDetail = () => {
  selectedGame.value = null;
};

const handleLaunch = async (gameTitle: string) => {
  try {
    await invoke("launch_game");
    console.log(`${gameTitle} を起動`);
  } catch (error) {
    console.error("起動失敗", error);
  }
};
</script>

<template>
  <div class="library-view">
    
    <div v-if="!selectedGame">
      <h2>ライブラリ</h2>
      <div class="game-list">
        <div 
          v-for="game in games" 
          :key="game.id" 
          class="game-item" 
          @click="showDetail(game)"
        >
          <div class="game-image-wrapper">
            <img :src="game.image" :alt="game.title" />
          </div>
          <div class="game-info">
            <div class="game-title">{{ game.title }}</div>
            <div class="game-explanation">{{ game.explanation }}</div>
          </div>
          <div class="game-action">
            <button class="launch-button" @click.stop="handleLaunch(game.title)">プレイ</button>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="game-detail-view">
      <button class="back-button" @click="closeDetail">← 一覧に戻る</button>
      
      <div class="detail-header">
        <h2>{{ selectedGame.title }}</h2>
        <button class="launch-button detail-launch" @click="handleLaunch(selectedGame.title)">
          プレイ
        </button>
      </div>

      <div class="detail-image-wrapper">
        <img :src="selectedGame.image" :alt="selectedGame.title" />
      </div>
      
      <div class="detail-content">
        <h3>概要</h3>
        <p>{{ selectedGame.detail }}</p>
      </div>
    </div>

  </div>
</template>

<style scoped>
/* 共通スタイル */
h2 {
  margin-top: 0;
  margin-bottom: 20px;
  font-size: 24px;
  border-left: 4px solid #ffff00;
  padding-left: 10px;
}
.launch-button { background-color: #ff8c00; color: white; border: none; padding: 10px 24px; font-size: 16px; font-weight: bold; border-radius: 4px; cursor: pointer; transition: background-color 0.2s; }
.launch-button:hover { background-color: #ffa500; }
.launch-button:active { background-color: #d2691e; }

/* 一覧画面スタイル */
.game-list { display: flex; flex-direction: column; gap: 15px; }
.game-item { 
  display: flex; 
  align-items: center; 
  background-color: #252526; 
  border: 1px solid #333333; 
  border-radius: 8px; 
  padding: 15px; 
  transition: transform 0.2s, background-color 0.2s;
  cursor: pointer; /* 押せることを明示 */
}
.game-item:hover { background-color: #2d2d30; transform: translateX(5px); }
.game-image-wrapper { width: 120px; height: 72px; border-radius: 4px; overflow: hidden; background-color: #000; flex-shrink: 0; }
.game-image-wrapper img { width: 100%; height: 100%; object-fit: cover; }
.game-info { flex: 1; padding: 0 20px; }
.game-title { font-size: 18px; font-weight: bold; margin-bottom: 5px; }
.game-explanation { font-size: 14px; color: #888888; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; max-width: 400px; }
.game-action { flex-shrink: 0; }

/* 詳細画面スタイル */
.game-detail-view {
  animation: fadeIn 0.3s ease;
}
@keyframes fadeIn {
  from { opacity: 0; transform: translateY(10px); }
  to { opacity: 1; transform: translateY(0); }
}
.back-button {
  background: none;
  border: none;
  color: #aaaaaa;
  font-size: 16px;
  cursor: pointer;
  padding: 0 0 20px 0;
  transition: color 0.2s;
}
.back-button:hover { color: #ffffff; }
.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}
.detail-header h2 { margin-bottom: 0; }
.detail-launch { font-size: 18px; padding: 12px 32px; }
.detail-image-wrapper {
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
  margin-bottom: 20px;
}
.detail-image-wrapper img {
  width: 100%;
  height: 100%;
  object-fit: cover;
}
.detail-content h3 {
  font-size: 18px;
  border-bottom: 1px solid #333;
  padding-bottom: 8px;
  margin-bottom: 10px;
}
.detail-content p {
  color: #cccccc;
  line-height: 1.6;
}
</style>