use std::process::Command;

#[tauri::command]
fn launch_game() -> Result<(), String> {

    Command::new("C:\\QpicGameLauncher\\games\\my_test_game.exe")
        .spawn()
        .map_err(|e| e.to_string())?;
    
    Ok(())
}

#[cfg_attr(mobile, tauri::mobile_entry_point)]
pub fn run() {
    tauri::Builder::default()
        .plugin(tauri_plugin_opener::init())
        .invoke_handler(tauri::generate_handler![launch_game])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
