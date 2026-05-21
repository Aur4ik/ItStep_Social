import { useEffect, useState } from "react";
import { useParams, useNavigate } from "react-router-dom";
import {
  getCommunityById, getCommunityPosts, getCommunityMembers,
  leaveCommunity, createCommunityPost,
} from "../api/communities";
import PostCard from "../components/PostCard";
import { useAuth } from "../context/AuthContext";

export default function CommunityDetailPage() {
  const { id } = useParams();
  const { user } = useAuth();
  const navigate = useNavigate();
  const [community, setCommunity] = useState(null);
  const [posts, setPosts] = useState([]);
  const [members, setMembers] = useState([]);
  const [content, setContent] = useState("");
  const [tab, setTab] = useState("posts");

  useEffect(() => {
    getCommunityById(id).then((res) => setCommunity(res.data)).catch(() => {});
    getCommunityPosts(id).then((res) => setPosts(res.data || [])).catch(() => {});
    getCommunityMembers(id).then((res) => setMembers(res.data || [])).catch(() => {});
  }, [id]);

  const isMember = members.some((m) => Number(m.id) === Number(user?.id));

  const handlePost = async (e) => {
    e.preventDefault();
    if (!content.trim()) return;
    const res = await createCommunityPost(id, { content });
    const postWithAuthor = { ...res.data, author: user };
    setPosts((prev) => [postWithAuthor, ...prev]);
    setContent("");
  };

  const handleLeave = async () => {
    await leaveCommunity(id);
    setMembers((prev) => prev.filter((m) => Number(m.id) !== Number(user?.id)));
  };

  if (!community) return (
    <div style={{ textAlign: "center", marginTop: 60, color: "#999" }}>Загрузка...</div>
  );

  return (
    <div style={styles.wrapper}>
      <div style={styles.header}>
        <h2>{community.name}</h2>
        <p style={styles.desc}>{community.description}</p>
        <p style={styles.membersCount}>👥 Участников: {community.members_count || members.length}</p>
        <div style={styles.headerActions}>
          <button onClick={() => navigate(`/communities/${id}/schedule`)} style={styles.scheduleBtn}>
            📅 Расписание
          </button>
          {isMember && (
            <button onClick={handleLeave} style={styles.leaveBtn}>Выйти</button>
          )}
        </div>
      </div>

      <div style={styles.tabs}>
        <button onClick={() => setTab("posts")} style={tab === "posts" ? styles.activeTab : styles.tab}>Посты</button>
        <button onClick={() => setTab("members")} style={tab === "members" ? styles.activeTab : styles.tab}>
          Участники ({members.length})
        </button>
      </div>

      {tab === "posts" && (
        <div>
          {isMember && (
            <form onSubmit={handlePost} style={styles.form}>
              <textarea
                value={content} onChange={(e) => setContent(e.target.value)}
                placeholder="Написать в сообщество..." style={styles.textarea}
              />
              <button type="submit" style={styles.btn}>Опубликовать</button>
            </form>
          )}
          {posts.map((p) => (
            <PostCard key={p.id} post={p} onDelete={(pid) => setPosts((prev) => prev.filter((x) => x.id !== pid))} />
          ))}
        </div>
      )}

      {tab === "members" && (
        <div style={styles.memberList}>
          {members.map((m) => (
            <div key={m.id} style={styles.member}>
              <strong>{m.first_name} {m.last_name}</strong>
              <span style={styles.memberRole}>{m.role}</span>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

const styles = {
  wrapper: { maxWidth: 640, margin: "24px auto", padding: "0 16px" },
  header: { background: "#fff", borderRadius: 10, padding: 20, marginBottom: 16, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  desc: { color: "#777", margin: "6px 0 4px" },
  membersCount: { color: "#999", fontSize: 13, margin: "0 0 12px" },
  headerActions: { display: "flex", gap: 8 },
  scheduleBtn: { background: "#4a90e2", color: "#fff", border: "none", borderRadius: 8, padding: "6px 16px", cursor: "pointer" },
  leaveBtn: { background: "#ff6b6b", color: "#fff", border: "none", borderRadius: 8, padding: "6px 16px", cursor: "pointer" },
  tabs: { display: "flex", gap: 8, marginBottom: 16 },
  tab: { padding: "8px 20px", border: "1px solid #ddd", borderRadius: 8, background: "#fff", cursor: "pointer" },
  activeTab: { padding: "8px 20px", border: "none", borderRadius: 8, background: "#e94560", color: "#fff", cursor: "pointer" },
  form: { background: "#fff", borderRadius: 10, padding: 16, marginBottom: 16, display: "flex", flexDirection: "column", gap: 10, boxShadow: "0 2px 8px rgba(0,0,0,0.08)" },
  textarea: { padding: 10, border: "1px solid #ddd", borderRadius: 8, minHeight: 70, resize: "vertical", fontSize: 14 },
  btn: { alignSelf: "flex-end", background: "#e94560", color: "#fff", border: "none", borderRadius: 8, padding: "8px 18px", cursor: "pointer" },
  memberList: { display: "flex", flexDirection: "column", gap: 8 },
  member: { background: "#fff", borderRadius: 8, padding: "10px 16px", display: "flex", justifyContent: "space-between", alignItems: "center", boxShadow: "0 1px 4px rgba(0,0,0,0.06)" },
  memberRole: { background: "#f0f0f0", borderRadius: 20, padding: "2px 10px", fontSize: 12 },
};