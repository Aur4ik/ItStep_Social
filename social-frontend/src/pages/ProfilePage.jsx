import { useRef } from "react";
import { useAuth } from "../context/AuthContext";
import { uploadAvatar, getMe } from "../api/auth";

export default function ProfilePage() {
  const { user, setUser } = useAuth();
  const fileRef = useRef();

  const handleAvatarChange = async (e) => {
    const file = e.target.files[0];
    if (!file) return;
    const formData = new FormData();
    formData.append("avatar", file);
    try {
      await uploadAvatar(formData);
      const res = await getMe();
      setUser(res.data);
    } catch {
      alert("Ошибка загрузки аватара");
    }
  };

  if (!user) return null;

  const avatarUrl = user.avatar
    ? `http://localhost:8080${user.avatar}`
    : null;

  return (
    <div style={styles.wrapper}>
      <div style={styles.card}>
        <div style={styles.avatarWrap} onClick={() => fileRef.current.click()}>
          {avatarUrl ? (
            <img src={avatarUrl} alt="avatar" style={styles.avatar} />
          ) : (
            <div style={styles.avatarPlaceholder}>👤</div>
          )}
          <div style={styles.editOverlay}>Изменить</div>
        </div>
        <input
          type="file" ref={fileRef} style={{ display: "none" }}
          accept="image/*" onChange={handleAvatarChange}
        />
        <h2>{user.first_name} {user.last_name}</h2>
        <p style={styles.info}>📧 {user.email}</p>
        <p style={styles.info}>🎓 Группа: {user.group || "—"}</p>
        <p style={styles.info}>
          🏷️ Роль: <span style={styles.role}>{user.role}</span>
        </p>
      </div>
    </div>
  );
}

const styles = {
  wrapper: { display: "flex", justifyContent: "center", padding: "40px 16px" },
  card: { background: "#fff", borderRadius: 12, padding: 32, width: 360, textAlign: "center", boxShadow: "0 4px 20px rgba(0,0,0,0.1)" },
  avatarWrap: { position: "relative", width: 100, height: 100, margin: "0 auto 20px", cursor: "pointer" },
  avatar: { width: 100, height: 100, borderRadius: "50%", objectFit: "cover" },
  avatarPlaceholder: { width: 100, height: 100, borderRadius: "50%", background: "#eee", display: "flex", alignItems: "center", justifyContent: "center", fontSize: 48 },
  editOverlay: { position: "absolute", bottom: 0, width: "100%", textAlign: "center", background: "rgba(0,0,0,0.4)", color: "#fff", fontSize: 12, borderRadius: "0 0 50px 50px", padding: "2px 0" },
  info: { color: "#555", margin: "6px 0" },
  role: { background: "#e94560", color: "#fff", padding: "2px 10px", borderRadius: 20, fontSize: 13 },
};