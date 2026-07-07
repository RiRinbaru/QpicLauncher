<script setup lang="ts">
import { ref, onMounted } from "vue";

interface Game {
  id: string;
  title: string;
  creatorId: string;
  genre: string;
  tags: string;
  explanation: string;
}

const games = ref<Game[]>([]);
const isLoading = ref<boolean>(true);
const errorMessage = ref<string>("");

const selectedGame = ref<Game | null>(null);

// データを取得
const fetchGames = async () => {
  try {
    const response = await fetch("http://localhost:8080/api/games");
    if (!response.ok) {
      throw new Error(`サーバーエラー: ${response.status}`);
    }
    const data = await response.json();
    games.value = data;
  } catch (error: any) {
    errorMessage.value = "サーバーとの通信に失敗した。\n" + error.message;
  } finally {
    isLoading.value = false;
  }
};

onMounted(() => {
  fetchGames();
});

// 詳細画面を開く
const showDetail = (game: Game) => {
  selectedGame.value = game;
};

// 一覧画面に戻る
const closeDetail = () => {
  selectedGame.value = null;
};

// ダウンロード処理（現状はモックアップ）
const handleDownload = (gameTitle: string) => {
  alert(`${gameTitle} のダウンロード処理をここに実装していく`);
};
</script>

<template>
  <div class="download-view">
    
    <div v-if="!selectedGame">
      <h2>ダウンロード</h2>

      <div v-if="isLoading" class="status-message">
        サーバーからゲーム一覧を取得中だ...
      </div>
      <div v-else-if="errorMessage" class="status-message error">
        {{ errorMessage }}
      </div>
      <div v-else class="game-list">
        <div 
          v-for="game in games" 
          :key="game.id" 
          class="game-item"
          @click="showDetail(game)"
        >
          <div class="game-info">
            <div class="game-title">{{ game.title }}</div>
            <div class="game-meta">
              <span>ジャンル: {{ game.genre }}</span> | <span>タグ: {{ game.tags }}</span>
            </div>
            <div class="game-explanation">{{ game.explanation }}</div>
          </div>
          <div class="game-action">
            <button class="download-button" @click.stop="showDetail(game)">
              詳細 / DL
            </button>
          </div>
        </div>
      </div>
    </div>

    <div v-else class="game-detail-view">
      <button class="back-button" @click="closeDetail">← 一覧に戻る</button>
      
      <div class="detail-header">
        <h2>{{ selectedGame.title }}</h2>
        <button class="download-button detail-download" @click="handleDownload(selectedGame.title)">
          ダウンロード
        </button>
      </div>

      <div class="detail-content">
        <h3>概要</h3>
        <p>{{ selectedGame.explanation }}</p>
        
        <h3>データ情報</h3>
        <ul>
          <li><strong>ジャンル:</strong> {{ selectedGame.genre }}</li>
          <li><strong>タグ:</strong> {{ selectedGame.tags }}</li>
          <li><strong>クリエイターID:</strong> {{ selectedGame.creatorId }}</li>
        </ul>
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
  border-left: 4px solid #00ff00; /* ライブラリと色を分けて区別する */
  padding-left: 10px;
}

.status-message {
  padding: 20px;
  background-color: #2d2d30;
  border-radius: 8px;
  text-align: center;
  color: #aaaaaa;
}
.status-message.error {
  color: #ff5555;
  border: 1px solid #ff5555;
}

.download-button { background-color: #28a745; color: white; border: none; padding: 10px 24px; font-size: 16px; font-weight: bold; border-radius: 4px; cursor: pointer; transition: background-color 0.2s; }
.download-button:hover { background-color: #218838; }
.download-button:active { background-color: #1e7e34; }

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
  cursor: pointer;
}
.game-item:hover { background-color: #2d2d30; transform: translateX(5px); }
.game-info { flex: 1; padding: 0 20px 0 0; }
.game-title { font-size: 18px; font-weight: bold; margin-bottom: 5px; }
.game-meta { font-size: 12px; color: #00ff00; margin-bottom: 8px; }
.game-explanation { font-size: 14px; color: #888888; overflow: hidden; text-overflow: ellipsis; white-space: nowrap; max-width: 400px;}
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
.detail-download { font-size: 18px; padding: 12px 32px; }
.detail-content h3 {
  font-size: 18px;
  border-bottom: 1px solid #333;
  padding-bottom: 8px;
  margin-bottom: 10px;
  margin-top: 20px;
}
.detail-content p, .detail-content ul {
  color: #cccccc;
  line-height: 1.6;
}
.detail-content ul {
  list-style-type: none;
  padding: 0;
}
.detail-content li {
  margin-bottom: 5px;
}
</style>