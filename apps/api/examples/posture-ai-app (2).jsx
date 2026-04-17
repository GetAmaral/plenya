import { useState, useRef, useEffect, useCallback } from "react";

// ─── Constants & Data ────────────────────────────────────────────
const KEYPOINT_NAMES = [
  "nariz","olho_esq","olho_dir","orelha_esq","orelha_dir",
  "ombro_esq","ombro_dir","cotovelo_esq","cotovelo_dir",
  "punho_esq","punho_dir","quadril_esq","quadril_dir",
  "joelho_esq","joelho_dir","tornozelo_esq","tornozelo_dir"
];

const KEYPOINT_LABELS = {
  nariz:"Nariz",olho_esq:"Olho Esq.",olho_dir:"Olho Dir.",
  orelha_esq:"Orelha Esq.",orelha_dir:"Orelha Dir.",
  ombro_esq:"Ombro Esq.",ombro_dir:"Ombro Dir.",
  cotovelo_esq:"Cotovelo Esq.",cotovelo_dir:"Cotovelo Dir.",
  punho_esq:"Punho Esq.",punho_dir:"Punho Dir.",
  quadril_esq:"Quadril Esq.",quadril_dir:"Quadril Dir.",
  joelho_esq:"Joelho Esq.",joelho_dir:"Joelho Dir.",
  tornozelo_esq:"Tornozelo Esq.",tornozelo_dir:"Tornozelo Dir."
};

const SKELETON_CONNECTIONS = [
  [5,6],[5,7],[7,9],[6,8],[8,10],[5,11],[6,12],[11,12],[11,13],[13,15],[12,14],[14,16]
];

const VIEW_TYPES = [
  { id:"front", label:"Frontal", icon:"👤" },
  { id:"side_left", label:"Lateral Esq.", icon:"🧍" },
  { id:"side_right", label:"Lateral Dir.", icon:"🧍" },
  { id:"back", label:"Posterior", icon:"👤" },
];

const POSTURE_TIPS = [
  "Mantenha os ombros relaxados e alinhados horizontalmente.",
  "A cabeça deve estar centralizada sobre os ombros, sem projeção anterior.",
  "Fortaleça o core para dar suporte à coluna lombar.",
  "Evite cruzar as pernas ao sentar — isso desalinha a pelve.",
  "Faça pausas a cada 30 min se trabalha sentado.",
  "Exercícios de mobilidade torácica ajudam a corrigir cifose.",
  "Alongue os flexores do quadril diariamente.",
  "Use um monitor na altura dos olhos para evitar flexão cervical.",
];

const VIDEO_ANGLES = [
  { id: "fhp", name: "Anteriorização (FHP)", points: [3, 5], type: "vertical", color: "#f59e0b" },
  { id: "trunk", name: "Tronco", points: [5, 11], type: "vertical", color: "#6366f1" },
  { id: "knee", name: "Joelho", points: [11, 13, 15], type: "angle", color: "#22d3ee" },
  { id: "hip", name: "Quadril", points: [5, 11, 13], type: "angle", color: "#f472b6" },
];

// ─── Utility Functions ───────────────────────────────────────────
function angleBetween(p1, p2, p3) {
  const a = { x: p1.x - p2.x, y: p1.y - p2.y };
  const b = { x: p3.x - p2.x, y: p3.y - p2.y };
  const dot = a.x * b.x + a.y * b.y;
  const magA = Math.sqrt(a.x * a.x + a.y * a.y);
  const magB = Math.sqrt(b.x * b.x + b.y * b.y);
  if (magA === 0 || magB === 0) return 0;
  return (Math.acos(Math.max(-1, Math.min(1, dot / (magA * magB)))) * 180) / Math.PI;
}

function angleTwoPoints(p1, p2) {
  return (Math.atan2(p2.y - p1.y, p2.x - p1.x) * 180) / Math.PI;
}

function verticalAngle(p1, p2) {
  return Math.abs(angleTwoPoints(p1, p2) + 90);
}

function distance(p1, p2) {
  return Math.sqrt((p2.x - p1.x) ** 2 + (p2.y - p1.y) ** 2);
}

function midpoint(p1, p2) {
  return { x: (p1.x + p2.x) / 2, y: (p1.y + p2.y) / 2 };
}

function classifySeverity(deviation, thresholds) {
  if (Math.abs(deviation) <= thresholds[0]) return { level: "Normal", color: "#22c55e", emoji: "✅" };
  if (Math.abs(deviation) <= thresholds[1]) return { level: "Leve", color: "#eab308", emoji: "⚠️" };
  if (Math.abs(deviation) <= thresholds[2]) return { level: "Moderado", color: "#f97316", emoji: "🔶" };
  return { level: "Severo", color: "#ef4444", emoji: "🔴" };
}

// ─── Posture Analysis Engine ─────────────────────────────────────
function analyzePosture(keypoints, viewType) {
  const kp = {};
  keypoints.forEach((p, i) => { kp[KEYPOINT_NAMES[i]] = p; });
  const results = [];

  if (viewType === "front" || viewType === "back") {
    if (kp.ombro_esq && kp.ombro_dir) {
      const angle = angleTwoPoints(kp.ombro_esq, kp.ombro_dir);
      const dev = Math.abs(angle);
      const sev = classifySeverity(dev, [2, 5, 10]);
      results.push({ name: "Alinhamento dos Ombros", value: `${dev.toFixed(1)}°`, deviation: dev, description: dev < 2 ? "Ombros nivelados" : `Inclinação de ${dev.toFixed(1)}°`, ...sev });
    }
    if (kp.quadril_esq && kp.quadril_dir) {
      const angle = angleTwoPoints(kp.quadril_esq, kp.quadril_dir);
      const dev = Math.abs(angle);
      const sev = classifySeverity(dev, [2, 4, 8]);
      results.push({ name: "Alinhamento Pélvico", value: `${dev.toFixed(1)}°`, deviation: dev, description: dev < 2 ? "Pelve nivelada" : `Inclinação pélvica de ${dev.toFixed(1)}°`, ...sev });
    }
    if (kp.olho_esq && kp.olho_dir) {
      const angle = angleTwoPoints(kp.olho_esq, kp.olho_dir);
      const dev = Math.abs(angle);
      const sev = classifySeverity(dev, [3, 6, 12]);
      results.push({ name: "Inclinação da Cabeça", value: `${dev.toFixed(1)}°`, deviation: dev, description: dev < 3 ? "Cabeça alinhada" : `Inclinação lateral de ${dev.toFixed(1)}°`, ...sev });
    }
    if (kp.ombro_esq && kp.ombro_dir && kp.quadril_esq && kp.quadril_dir) {
      const shoulderMid = midpoint(kp.ombro_esq, kp.ombro_dir);
      const hipMid = midpoint(kp.quadril_esq, kp.quadril_dir);
      const trunkAngle = Math.abs(angleTwoPoints(hipMid, shoulderMid) + 90);
      const sev = classifySeverity(trunkAngle, [2, 5, 10]);
      results.push({ name: "Simetria do Tronco", value: `${trunkAngle.toFixed(1)}°`, deviation: trunkAngle, description: trunkAngle < 2 ? "Tronco simétrico" : `Desvio lateral de ${trunkAngle.toFixed(1)}°`, ...sev });
    }
    if (kp.quadril_esq && kp.joelho_esq && kp.tornozelo_esq) {
      const angle = angleBetween(kp.quadril_esq, kp.joelho_esq, kp.tornozelo_esq);
      const dev = Math.abs(180 - angle);
      const sev = classifySeverity(dev, [5, 10, 18]);
      results.push({ name: "Joelho Esquerdo (Valgo/Varo)", value: `${angle.toFixed(1)}°`, deviation: dev, description: dev < 5 ? "Alinhamento normal" : angle < 175 ? `Tendência a valgo (${dev.toFixed(1)}°)` : `Tendência a varo (${dev.toFixed(1)}°)`, ...sev });
    }
    if (kp.quadril_dir && kp.joelho_dir && kp.tornozelo_dir) {
      const angle = angleBetween(kp.quadril_dir, kp.joelho_dir, kp.tornozelo_dir);
      const dev = Math.abs(180 - angle);
      const sev = classifySeverity(dev, [5, 10, 18]);
      results.push({ name: "Joelho Direito (Valgo/Varo)", value: `${angle.toFixed(1)}°`, deviation: dev, description: dev < 5 ? "Alinhamento normal" : angle < 175 ? `Tendência a valgo (${dev.toFixed(1)}°)` : `Tendência a varo (${dev.toFixed(1)}°)`, ...sev });
    }
    if (kp.olho_esq && kp.olho_dir && kp.nariz) {
      const dLeft = distance(kp.olho_esq, kp.nariz);
      const dRight = distance(kp.olho_dir, kp.nariz);
      const ratio = Math.min(dLeft, dRight) / Math.max(dLeft, dRight) * 100;
      const dev = 100 - ratio;
      const sev = classifySeverity(dev, [3, 8, 15]);
      results.push({ name: "Simetria Facial", value: `${ratio.toFixed(1)}%`, deviation: dev, description: dev < 3 ? "Alta simetria facial" : `Assimetria de ${dev.toFixed(1)}%`, ...sev });
    }
  }

  if (viewType === "side_left" || viewType === "side_right") {
    if (kp.orelha_esq && kp.ombro_esq) {
      const ref = viewType === "side_left" ? kp.orelha_esq : kp.orelha_dir;
      const shoulder = viewType === "side_left" ? kp.ombro_esq : kp.ombro_dir;
      const deviation = verticalAngle(shoulder, ref);
      const sev = classifySeverity(deviation, [5, 15, 25]);
      results.push({ name: "Anteriorização da Cabeça (FHP)", value: `${deviation.toFixed(1)}°`, deviation, description: deviation < 5 ? "Cabeça bem alinhada" : `Projeção anterior de ${deviation.toFixed(1)}°`, ...sev });
    }
    if (kp.ombro_esq && kp.quadril_esq) {
      const s = viewType === "side_left" ? kp.ombro_esq : kp.ombro_dir;
      const h = viewType === "side_left" ? kp.quadril_esq : kp.quadril_dir;
      const deviation = verticalAngle(h, s);
      const sev = classifySeverity(deviation, [5, 12, 20]);
      results.push({ name: "Alinhamento do Tronco", value: `${deviation.toFixed(1)}°`, deviation, description: deviation < 5 ? "Coluna bem alinhada" : `Desvio de ${deviation.toFixed(1)}°`, ...sev });
    }
    if (kp.quadril_esq && kp.joelho_esq && kp.tornozelo_esq) {
      const hip = viewType === "side_left" ? kp.quadril_esq : kp.quadril_dir;
      const knee = viewType === "side_left" ? kp.joelho_esq : kp.joelho_dir;
      const ankle = viewType === "side_left" ? kp.tornozelo_esq : kp.tornozelo_dir;
      const angle = angleBetween(hip, knee, ankle);
      const dev = Math.abs(180 - angle);
      const sev = classifySeverity(dev, [5, 10, 20]);
      results.push({ name: "Flexão/Hiperextensão do Joelho", value: `${angle.toFixed(1)}°`, deviation: dev, description: dev < 5 ? "Joelho em posição neutra" : angle < 175 ? `Flexão de ${dev.toFixed(1)}°` : `Hiperextensão de ${dev.toFixed(1)}°`, ...sev });
    }
  }

  if (kp.nariz && kp.ombro_esq && kp.quadril_esq && kp.joelho_esq && kp.tornozelo_esq) {
    const headToShoulder = distance(kp.nariz, midpoint(kp.ombro_esq, kp.ombro_dir || kp.ombro_esq));
    const shoulderToHip = distance(midpoint(kp.ombro_esq, kp.ombro_dir || kp.ombro_esq), midpoint(kp.quadril_esq, kp.quadril_dir || kp.quadril_esq));
    const hipToKnee = distance(midpoint(kp.quadril_esq, kp.quadril_dir || kp.quadril_esq), midpoint(kp.joelho_esq, kp.joelho_dir || kp.joelho_esq));
    const kneeToAnkle = distance(midpoint(kp.joelho_esq, kp.joelho_dir || kp.joelho_esq), midpoint(kp.tornozelo_esq, kp.tornozelo_dir || kp.tornozelo_esq));
    const totalHeight = headToShoulder + shoulderToHip + hipToKnee + kneeToAnkle;
    const navelRatio = (hipToKnee + kneeToAnkle) / totalHeight;
    const goldenDev = Math.abs(navelRatio - 0.618) * 100;
    const sev = classifySeverity(goldenDev, [3, 8, 15]);
    results.push({ name: "Proporção Áurea", value: `${navelRatio.toFixed(3)}`, deviation: goldenDev, description: `Razão: ${navelRatio.toFixed(3)} (ideal: 0.618)`, ...sev });
  }

  if (results.length > 0) {
    const penalties = results.reduce((sum, r) => sum + Math.min(r.deviation * 1.5, 25), 0);
    const score = Math.max(0, Math.round(100 - penalties));
    return { results, score };
  }
  return { results, score: null };
}

// ─── TF.js Loader ────────────────────────────────────────────────
function loadScript(src) {
  return new Promise((resolve, reject) => {
    if (document.querySelector(`script[src="${src}"]`)) { resolve(); return; }
    const s = document.createElement("script");
    s.src = src; s.onload = resolve; s.onerror = reject;
    document.head.appendChild(s);
  });
}

async function loadPoseDetector() {
  if (typeof window.tf === "undefined") {
    await loadScript("https://cdn.jsdelivr.net/npm/@tensorflow/tfjs-core@4.17.0/dist/tf-core.min.js");
    await loadScript("https://cdn.jsdelivr.net/npm/@tensorflow/tfjs-converter@4.17.0/dist/tf-converter.min.js");
    await loadScript("https://cdn.jsdelivr.net/npm/@tensorflow/tfjs-backend-webgl@4.17.0/dist/tf-backend-webgl.min.js");
    await loadScript("https://cdn.jsdelivr.net/npm/@tensorflow-models/pose-detection@2.1.3/dist/pose-detection.min.js");
  }
  await window.tf.setBackend("webgl");
  await window.tf.ready();
  const detector = await window.poseDetection.createDetector(
    window.poseDetection.SupportedModels.MoveNet,
    { modelType: window.poseDetection.movenet.modelType.SINGLEPOSE_THUNDER }
  );
  return detector;
}

// ─── Simple Line Chart Component ─────────────────────────────────
function MiniChart({ data, colors, labels, height = 160 }) {
  if (!data || data.length === 0 || !data[0] || data[0].length === 0) return null;
  const maxVal = Math.max(...data.flat(), 1);
  const minVal = Math.min(...data.flat(), 0);
  const range = maxVal - minVal || 1;
  const w = 100;
  const h = height;
  const pad = 4;

  return (
    <div style={{ width: "100%", overflowX: "auto" }}>
      <svg viewBox={`0 0 ${w} ${h}`} width="100%" height={height} preserveAspectRatio="none" style={{ display: "block" }}>
        <rect x="0" y="0" width={w} height={h} fill="rgba(15,15,30,0.5)" rx="4" />
        {[0, 0.25, 0.5, 0.75, 1].map((f, i) => (
          <g key={i}>
            <line x1={pad} y1={pad + (h - 2 * pad) * f} x2={w - pad} y2={pad + (h - 2 * pad) * f} stroke="rgba(100,116,139,0.15)" strokeWidth="0.3" />
            <text x={pad + 1} y={pad + (h - 2 * pad) * f - 1} fill="#64748b" fontSize="3" fontFamily="sans-serif">
              {(maxVal - range * f).toFixed(0)}°
            </text>
          </g>
        ))}
        {data.map((series, si) => {
          if (series.length < 2) return null;
          const points = series.map((val, i) => {
            const x = pad + (i / (series.length - 1)) * (w - 2 * pad);
            const y = pad + ((maxVal - val) / range) * (h - 2 * pad);
            return `${x},${y}`;
          }).join(" ");
          return <polyline key={si} points={points} fill="none" stroke={colors[si]} strokeWidth="0.8" strokeLinejoin="round" />;
        })}
      </svg>
      <div style={{ display: "flex", gap: 12, justifyContent: "center", marginTop: 6, flexWrap: "wrap" }}>
        {labels.map((label, i) => (
          <div key={i} style={{ display: "flex", alignItems: "center", gap: 4 }}>
            <div style={{ width: 10, height: 3, background: colors[i], borderRadius: 2 }} />
            <span style={{ fontSize: 10, color: "#94a3b8" }}>{label}</span>
          </div>
        ))}
      </div>
    </div>
  );
}

// ─── Main App Component ──────────────────────────────────────────
export default function PostureApp() {
  const [screen, setScreen] = useState("home");
  const [image, setImage] = useState(null);
  const [imgDims, setImgDims] = useState({ w: 0, h: 0 });
  const [keypoints, setKeypoints] = useState([]);
  const [selectedKp, setSelectedKp] = useState(null);
  const [dragging, setDragging] = useState(null);
  const [viewType, setViewType] = useState("front");
  const [analysis, setAnalysis] = useState(null);
  const [loading, setLoading] = useState(false);
  const [loadingMsg, setLoadingMsg] = useState("");
  const [detectorReady, setDetectorReady] = useState(false);
  const [showSkeleton, setShowSkeleton] = useState(true);
  const [history, setHistory] = useState([]);
  const [privacyMask, setPrivacyMask] = useState(false);
  const [customAngles, setCustomAngles] = useState([]);
  const [romMode, setRomMode] = useState(false);
  const [romPoints, setRomPoints] = useState([]);
  const [tipIndex, setTipIndex] = useState(0);

  // Video analysis state
  const [videoSrc, setVideoSrc] = useState(null);
  const [videoAnalysis, setVideoAnalysis] = useState(null);
  const [videoProcessing, setVideoProcessing] = useState(false);
  const [videoProgress, setVideoProgress] = useState(0);
  const [videoFrameIdx, setVideoFrameIdx] = useState(0);
  const [videoPlaying, setVideoPlaying] = useState(false);
  const [videoFrames, setVideoFrames] = useState([]);
  const [selectedVideoAngles, setSelectedVideoAngles] = useState(["fhp", "knee"]);

  const canvasRef = useRef(null);
  const imgRef = useRef(null);
  const detectorRef = useRef(null);
  const fileInputRef = useRef(null);
  const videoInputRef = useRef(null);
  const videoCanvasRef = useRef(null);
  const videoElRef = useRef(null);
  const playIntervalRef = useRef(null);

  useEffect(() => {
    const t = setInterval(() => setTipIndex(i => (i + 1) % POSTURE_TIPS.length), 6000);
    return () => clearInterval(t);
  }, []);

  useEffect(() => {
    loadPoseDetector().then(d => { detectorRef.current = d; setDetectorReady(true); }).catch(console.error);
  }, []);

  // Draw canvas for photo editor
  useEffect(() => {
    if (!canvasRef.current || !image || screen !== "editor") return;
    const canvas = canvasRef.current;
    const ctx = canvas.getContext("2d");
    const img = imgRef.current;
    if (!img) return;
    canvas.width = imgDims.w;
    canvas.height = imgDims.h;
    ctx.clearRect(0, 0, canvas.width, canvas.height);
    ctx.drawImage(img, 0, 0, imgDims.w, imgDims.h);

    if (privacyMask && keypoints.length > 0) {
      const nose = keypoints[0];
      const earL = keypoints[3];
      const earR = keypoints[4];
      if (nose && (earL || earR)) {
        const ref = earL || earR;
        const faceRadius = distance(nose, ref) * 1.4;
        ctx.save();
        ctx.beginPath();
        ctx.arc(nose.x, nose.y, faceRadius, 0, Math.PI * 2);
        ctx.fillStyle = "rgba(20,20,30,0.92)";
        ctx.fill();
        ctx.strokeStyle = "#6366f1";
        ctx.lineWidth = 2;
        ctx.stroke();
        ctx.fillStyle = "#a5b4fc";
        ctx.font = `${faceRadius * 0.5}px sans-serif`;
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillText("🔒", nose.x, nose.y);
        ctx.restore();
      }
    }

    if (showSkeleton && keypoints.length > 0) {
      ctx.save();
      SKELETON_CONNECTIONS.forEach(([i, j]) => {
        if (keypoints[i] && keypoints[j]) {
          ctx.beginPath();
          ctx.moveTo(keypoints[i].x, keypoints[i].y);
          ctx.lineTo(keypoints[j].x, keypoints[j].y);
          ctx.strokeStyle = "rgba(99,102,241,0.6)";
          ctx.lineWidth = 2.5;
          ctx.stroke();
        }
      });
      ctx.restore();
    }

    keypoints.forEach((kp, i) => {
      if (!kp) return;
      const isSelected = selectedKp === i;
      const r = isSelected ? 9 : 6;
      ctx.beginPath();
      ctx.arc(kp.x, kp.y, r, 0, Math.PI * 2);
      ctx.fillStyle = isSelected ? "#f59e0b" : kp.confidence > 0.3 ? "#22d3ee" : "#f87171";
      ctx.fill();
      ctx.strokeStyle = "#fff";
      ctx.lineWidth = 1.5;
      ctx.stroke();
      if (isSelected) {
        ctx.font = "bold 11px sans-serif";
        ctx.fillStyle = "#fff";
        ctx.textAlign = "center";
        ctx.fillText(KEYPOINT_LABELS[KEYPOINT_NAMES[i]], kp.x, kp.y - 14);
      }
    });

    if (romMode && romPoints.length >= 2) {
      ctx.save();
      for (let i = 0; i < romPoints.length - 1; i++) {
        ctx.beginPath();
        ctx.moveTo(romPoints[i].x, romPoints[i].y);
        ctx.lineTo(romPoints[i + 1].x, romPoints[i + 1].y);
        ctx.strokeStyle = "#f59e0b";
        ctx.lineWidth = 2;
        ctx.setLineDash([6, 4]);
        ctx.stroke();
      }
      ctx.setLineDash([]);
      romPoints.forEach((p, i) => {
        ctx.beginPath();
        ctx.arc(p.x, p.y, 7, 0, Math.PI * 2);
        ctx.fillStyle = "#f59e0b";
        ctx.fill();
        ctx.fillStyle = "#000";
        ctx.font = "bold 10px sans-serif";
        ctx.textAlign = "center";
        ctx.textBaseline = "middle";
        ctx.fillText(i + 1, p.x, p.y);
      });
      if (romPoints.length === 3) {
        const angle = angleBetween(romPoints[0], romPoints[1], romPoints[2]);
        ctx.font = "bold 14px sans-serif";
        ctx.fillStyle = "#f59e0b";
        ctx.textAlign = "center";
        ctx.fillText(`${angle.toFixed(1)}°`, romPoints[1].x, romPoints[1].y - 18);
      }
      ctx.restore();
    }
  }, [image, keypoints, selectedKp, showSkeleton, privacyMask, imgDims, romMode, romPoints, screen]);

  // Draw video frame overlay
  useEffect(() => {
    if (screen !== "video_results" || !videoCanvasRef.current || !videoFrames.length) return;
    const frame = videoFrames[videoFrameIdx];
    if (!frame) return;
    const canvas = videoCanvasRef.current;
    const ctx = canvas.getContext("2d");
    const w = frame.width;
    const h = frame.height;
    canvas.width = w;
    canvas.height = h;

    // Draw frame image
    if (frame.imageData) {
      ctx.putImageData(frame.imageData, 0, 0);
    } else {
      ctx.fillStyle = "#111";
      ctx.fillRect(0, 0, w, h);
    }

    // Draw skeleton
    const kps = frame.keypoints;
    if (kps) {
      ctx.save();
      SKELETON_CONNECTIONS.forEach(([i, j]) => {
        if (kps[i] && kps[j] && kps[i].confidence > 0.25 && kps[j].confidence > 0.25) {
          ctx.beginPath();
          ctx.moveTo(kps[i].x, kps[i].y);
          ctx.lineTo(kps[j].x, kps[j].y);
          ctx.strokeStyle = "rgba(99,102,241,0.7)";
          ctx.lineWidth = 2;
          ctx.stroke();
        }
      });
      kps.forEach((kp) => {
        if (!kp || kp.confidence < 0.25) return;
        ctx.beginPath();
        ctx.arc(kp.x, kp.y, 4, 0, Math.PI * 2);
        ctx.fillStyle = "#22d3ee";
        ctx.fill();
        ctx.strokeStyle = "#fff";
        ctx.lineWidth = 1;
        ctx.stroke();
      });

      // Draw selected angle arcs
      selectedVideoAngles.forEach(aid => {
        const angleDef = VIDEO_ANGLES.find(a => a.id === aid);
        if (!angleDef) return;
        const pts = angleDef.points.map(idx => kps[idx]).filter(p => p && p.confidence > 0.25);
        if (angleDef.type === "angle" && pts.length === 3) {
          ctx.beginPath();
          ctx.moveTo(pts[0].x, pts[0].y);
          ctx.lineTo(pts[1].x, pts[1].y);
          ctx.lineTo(pts[2].x, pts[2].y);
          ctx.strokeStyle = angleDef.color;
          ctx.lineWidth = 2.5;
          ctx.stroke();
          const angle = angleBetween(pts[0], pts[1], pts[2]);
          ctx.font = "bold 12px sans-serif";
          ctx.fillStyle = angleDef.color;
          ctx.textAlign = "center";
          ctx.fillText(`${angle.toFixed(0)}°`, pts[1].x + 15, pts[1].y - 8);
        } else if (angleDef.type === "vertical" && pts.length === 2) {
          ctx.beginPath();
          ctx.moveTo(pts[0].x, pts[0].y);
          ctx.lineTo(pts[1].x, pts[1].y);
          ctx.strokeStyle = angleDef.color;
          ctx.lineWidth = 2.5;
          ctx.stroke();
          // Vertical reference line
          ctx.beginPath();
          ctx.moveTo(pts[1].x, pts[1].y);
          ctx.lineTo(pts[1].x, pts[0].y);
          ctx.strokeStyle = "rgba(255,255,255,0.3)";
          ctx.lineWidth = 1;
          ctx.setLineDash([4, 4]);
          ctx.stroke();
          ctx.setLineDash([]);
          const angle = verticalAngle(pts[1], pts[0]);
          ctx.font = "bold 12px sans-serif";
          ctx.fillStyle = angleDef.color;
          ctx.textAlign = "center";
          ctx.fillText(`${angle.toFixed(0)}°`, (pts[0].x + pts[1].x) / 2 + 15, (pts[0].y + pts[1].y) / 2);
        }
      });

      ctx.restore();
    }
  }, [videoFrameIdx, videoFrames, screen, selectedVideoAngles]);

  // Video playback
  useEffect(() => {
    if (videoPlaying && videoFrames.length > 0) {
      playIntervalRef.current = setInterval(() => {
        setVideoFrameIdx(prev => {
          if (prev >= videoFrames.length - 1) {
            setVideoPlaying(false);
            return 0;
          }
          return prev + 1;
        });
      }, 120);
    }
    return () => clearInterval(playIntervalRef.current);
  }, [videoPlaying, videoFrames.length]);

  const handleImageUpload = (e) => {
    const file = e.target.files[0];
    if (!file) return;
    const reader = new FileReader();
    reader.onload = (ev) => {
      const img = new Image();
      img.onload = () => {
        const maxW = Math.min(600, window.innerWidth - 40);
        const scale = maxW / img.width;
        setImgDims({ w: maxW, h: img.height * scale });
        setImage(ev.target.result);
        setKeypoints([]);
        setAnalysis(null);
        setRomPoints([]);
        setCustomAngles([]);
        imgRef.current = img;
        setScreen("editor");
      };
      img.src = ev.target.result;
    };
    reader.readAsDataURL(file);
  };

  const handleVideoUpload = (e) => {
    const file = e.target.files[0];
    if (!file) return;
    const url = URL.createObjectURL(file);
    setVideoSrc(url);
    setVideoAnalysis(null);
    setVideoFrames([]);
    setVideoFrameIdx(0);
    setScreen("video_processing");
  };

  const processVideo = useCallback(async () => {
    if (!videoSrc || !detectorRef.current) return;
    setVideoProcessing(true);
    setVideoProgress(0);

    const video = document.createElement("video");
    video.src = videoSrc;
    video.muted = true;
    video.playsInline = true;

    await new Promise((resolve) => {
      video.onloadeddata = resolve;
      video.load();
    });

    const duration = video.duration;
    const fps = 5; // sample at 5fps for performance
    const totalFrames = Math.min(Math.floor(duration * fps), 150); // max 150 frames
    const maxW = Math.min(400, window.innerWidth - 40);
    const scale = maxW / video.videoWidth;
    const w = maxW;
    const h = Math.round(video.videoHeight * scale);

    const offCanvas = document.createElement("canvas");
    offCanvas.width = w;
    offCanvas.height = h;
    const offCtx = offCanvas.getContext("2d");

    const frames = [];
    const angleTimelines = {};
    VIDEO_ANGLES.forEach(a => { angleTimelines[a.id] = []; });

    for (let i = 0; i < totalFrames; i++) {
      const time = (i / fps);
      video.currentTime = time;
      await new Promise(r => { video.onseeked = r; });

      offCtx.drawImage(video, 0, 0, w, h);
      const imageData = offCtx.getImageData(0, 0, w, h);

      let kps = null;
      try {
        const poses = await detectorRef.current.estimatePoses(offCanvas);
        if (poses.length > 0) {
          kps = poses[0].keypoints.map(kp => ({
            x: kp.x, y: kp.y, confidence: kp.score
          }));

          // Calculate angles for timelines
          VIDEO_ANGLES.forEach(angleDef => {
            const pts = angleDef.points.map(idx => kps[idx]).filter(p => p && p.confidence > 0.25);
            let val = null;
            if (angleDef.type === "angle" && pts.length === 3) {
              val = angleBetween(pts[0], pts[1], pts[2]);
            } else if (angleDef.type === "vertical" && pts.length === 2) {
              val = verticalAngle(pts[1], pts[0]);
            }
            angleTimelines[angleDef.id].push(val !== null ? val : angleTimelines[angleDef.id].length > 0 ? angleTimelines[angleDef.id][angleTimelines[angleDef.id].length - 1] : 0);
          });
        } else {
          VIDEO_ANGLES.forEach(a => {
            angleTimelines[a.id].push(angleTimelines[a.id].length > 0 ? angleTimelines[a.id][angleTimelines[a.id].length - 1] : 0);
          });
        }
      } catch (err) {
        VIDEO_ANGLES.forEach(a => {
          angleTimelines[a.id].push(angleTimelines[a.id].length > 0 ? angleTimelines[a.id][angleTimelines[a.id].length - 1] : 0);
        });
      }

      frames.push({ keypoints: kps, imageData, width: w, height: h, time });
      setVideoProgress(Math.round(((i + 1) / totalFrames) * 100));
    }

    // Compute stats
    const stats = {};
    VIDEO_ANGLES.forEach(angleDef => {
      const vals = angleTimelines[angleDef.id].filter(v => v > 0);
      if (vals.length > 0) {
        const avg = vals.reduce((a, b) => a + b, 0) / vals.length;
        const max = Math.max(...vals);
        const min = Math.min(...vals);
        const range = max - min;
        stats[angleDef.id] = { avg, max, min, range, timeline: angleTimelines[angleDef.id] };
      }
    });

    setVideoFrames(frames);
    setVideoAnalysis({ stats, angleTimelines, totalFrames, duration, fps });
    setVideoProcessing(false);
    setScreen("video_results");
  }, [videoSrc]);

  // Auto-start video processing
  useEffect(() => {
    if (screen === "video_processing" && videoSrc && detectorReady && !videoProcessing) {
      processVideo();
    }
  }, [screen, videoSrc, detectorReady, videoProcessing, processVideo]);

  const runDetection = async () => {
    if (!detectorRef.current || !imgRef.current) return;
    setLoading(true);
    setLoadingMsg("Detectando pontos anatômicos com IA...");
    try {
      const tempCanvas = document.createElement("canvas");
      tempCanvas.width = imgDims.w;
      tempCanvas.height = imgDims.h;
      tempCanvas.getContext("2d").drawImage(imgRef.current, 0, 0, imgDims.w, imgDims.h);
      const poses = await detectorRef.current.estimatePoses(tempCanvas);
      if (poses.length > 0) {
        setKeypoints(poses[0].keypoints.map(kp => ({ x: kp.x, y: kp.y, confidence: kp.score })));
      }
    } catch (err) { console.error(err); }
    setLoading(false);
    setLoadingMsg("");
  };

  const runAnalysis = () => {
    const result = analyzePosture(keypoints, viewType);
    setAnalysis(result);
    setHistory(prev => [...prev, { date: new Date().toLocaleString("pt-BR"), viewType, score: result.score, results: result.results }]);
    setScreen("results");
  };

  const getCanvasPos = (e) => {
    const canvas = canvasRef.current;
    const rect = canvas.getBoundingClientRect();
    const scaleX = imgDims.w / rect.width;
    const scaleY = imgDims.h / rect.height;
    const clientX = e.touches ? e.touches[0].clientX : e.clientX;
    const clientY = e.touches ? e.touches[0].clientY : e.clientY;
    return { x: (clientX - rect.left) * scaleX, y: (clientY - rect.top) * scaleY };
  };

  const handleCanvasPointerDown = (e) => {
    const pos = getCanvasPos(e);
    if (romMode) {
      if (romPoints.length < 3) {
        const newPts = [...romPoints, pos];
        setRomPoints(newPts);
        if (newPts.length === 3) {
          setCustomAngles(prev => [...prev, { points: newPts, angle: angleBetween(newPts[0], newPts[1], newPts[2]) }]);
        }
      } else {
        setRomPoints([pos]);
      }
      return;
    }
    let closest = -1, minDist = 30;
    keypoints.forEach((kp, i) => {
      if (!kp) return;
      const d = distance(kp, pos);
      if (d < minDist) { minDist = d; closest = i; }
    });
    if (closest >= 0) { setDragging(closest); setSelectedKp(closest); }
    else setSelectedKp(null);
  };

  const handleCanvasPointerMove = (e) => {
    if (dragging === null) return;
    e.preventDefault();
    const pos = getCanvasPos(e);
    setKeypoints(prev => { const next = [...prev]; next[dragging] = { ...next[dragging], ...pos }; return next; });
  };

  const handleCanvasPointerUp = () => setDragging(null);

  const overallColor = (score) => score >= 80 ? "#22c55e" : score >= 60 ? "#eab308" : score >= 40 ? "#f97316" : "#ef4444";
  const overallLabel = (score) => score >= 80 ? "Excelente" : score >= 60 ? "Boa" : score >= 40 ? "Regular" : "Necessita Atenção";

  // ─── HOME ──────────────────────────────────────────────────────
  if (screen === "home") {
    return (
      <div style={S.app}>
        <div style={S.homeContainer}>
          <div style={S.logoArea}>
            <div style={S.logoIcon}>
              <svg viewBox="0 0 64 64" width="64" height="64">
                <circle cx="32" cy="12" r="8" fill="#6366f1" />
                <line x1="32" y1="20" x2="32" y2="44" stroke="#6366f1" strokeWidth="3" strokeLinecap="round"/>
                <line x1="32" y1="26" x2="18" y2="36" stroke="#6366f1" strokeWidth="3" strokeLinecap="round"/>
                <line x1="32" y1="26" x2="46" y2="36" stroke="#6366f1" strokeWidth="3" strokeLinecap="round"/>
                <line x1="32" y1="44" x2="22" y2="60" stroke="#6366f1" strokeWidth="3" strokeLinecap="round"/>
                <line x1="32" y1="44" x2="42" y2="60" stroke="#6366f1" strokeWidth="3" strokeLinecap="round"/>
                <circle cx="32" cy="12" r="4" fill="#22d3ee"/><circle cx="32" cy="26" r="3" fill="#22d3ee"/><circle cx="32" cy="44" r="3" fill="#22d3ee"/>
              </svg>
            </div>
            <h1 style={S.logoTitle}>PosturaAI</h1>
            <p style={S.logoSub}>Análise Postural com Inteligência Artificial</p>
          </div>
          <div style={S.tipCard}>
            <div style={{fontSize:22,flexShrink:0}}>💡</div>
            <p style={S.tipText}>{POSTURE_TIPS[tipIndex]}</p>
          </div>
          <div style={S.homeActions}>
            <button style={S.primaryBtn} onClick={() => fileInputRef.current?.click()}>
              <span style={{fontSize:22}}>📸</span><span>Avaliação por Foto</span>
            </button>
            <input ref={fileInputRef} type="file" accept="image/*" capture="environment" style={{display:"none"}} onChange={handleImageUpload}/>

            <button style={{...S.primaryBtn, background: "linear-gradient(135deg, #7c3aed, #6d28d9)"}} onClick={() => videoInputRef.current?.click()}>
              <span style={{fontSize:22}}>🎥</span><span>Análise de Vídeo</span>
            </button>
            <input ref={videoInputRef} type="file" accept="video/*" capture="environment" style={{display:"none"}} onChange={handleVideoUpload}/>

            <div style={S.viewSelector}>
              <p style={S.viewLabel}>Vista da Foto:</p>
              <div style={S.viewGrid}>
                {VIEW_TYPES.map(v => (
                  <button key={v.id} style={{...S.viewBtn, ...(viewType === v.id ? S.viewBtnActive : {})}} onClick={() => setViewType(v.id)}>
                    <span>{v.icon}</span><span style={{fontSize:11}}>{v.label}</span>
                  </button>
                ))}
              </div>
            </div>
            {history.length > 0 && (
              <button style={S.secondaryBtn} onClick={() => setScreen("history")}>
                <span>📊</span><span>Histórico ({history.length})</span>
              </button>
            )}
          </div>
          <div style={S.features}>
            {[
              ["🤖","Detecção Automática","MoveNet AI — 17 pontos anatômicos"],
              ["📐","Análise de Ângulos","Ombros, quadril, joelhos, cabeça e mais"],
              ["🎥","Análise de Vídeo","Postura dinâmica frame a frame com gráficos"],
              ["📏","Range of Motion","Meça ângulos customizados livremente"],
            ].map(([icon, title, desc], i) => (
              <div key={i} style={S.featureCard}>
                <span style={{fontSize:24}}>{icon}</span>
                <div><p style={S.featureTitle}>{title}</p><p style={S.featureDesc}>{desc}</p></div>
              </div>
            ))}
          </div>
          <p style={S.disclaimer}>⚕️ PosturaAI é uma ferramenta auxiliar. Resultados devem ser confirmados por profissional de saúde.</p>
          {!detectorReady && (
            <div style={S.loadingBar}>
              <div style={S.loadingPulse}/><span style={{fontSize:12,color:"#94a3b8"}}>Carregando modelo de IA...</span>
            </div>
          )}
        </div>
      </div>
    );
  }

  // ─── PHOTO EDITOR ──────────────────────────────────────────────
  if (screen === "editor") {
    const hasKp = keypoints.length > 0;
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("home")}>← Voltar</button>
          <h2 style={S.headerTitle}>{VIEW_TYPES.find(v => v.id === viewType)?.icon} {VIEW_TYPES.find(v => v.id === viewType)?.label}</h2>
          <div style={{width:60}}/>
        </div>
        <div style={S.canvasContainer}>
          {loading && <div style={S.loadingOverlay}><div style={S.spinner}/><p style={{color:"#e2e8f0",marginTop:12,fontSize:14}}>{loadingMsg}</p></div>}
          <canvas ref={canvasRef} style={{...S.canvas, width: imgDims.w, height: imgDims.h, maxWidth:"100%", cursor: romMode ? "crosshair" : dragging !== null ? "grabbing" : "grab"}}
            onMouseDown={handleCanvasPointerDown} onMouseMove={handleCanvasPointerMove} onMouseUp={handleCanvasPointerUp} onMouseLeave={handleCanvasPointerUp}
            onTouchStart={handleCanvasPointerDown} onTouchMove={handleCanvasPointerMove} onTouchEnd={handleCanvasPointerUp}/>
        </div>
        <div style={S.tools}>
          <div style={S.toolRow}>
            <button style={{...S.toolBtn, opacity: detectorReady ? 1 : 0.5}} onClick={runDetection} disabled={!detectorReady||loading}>🤖 {hasKp?"Re-detectar":"Detectar AI"}</button>
            {hasKp && <button style={{...S.toolBtn,...S.toolBtnAccent}} onClick={runAnalysis}>📊 Analisar</button>}
          </div>
          <div style={S.toolRow}>
            <button style={{...S.chip,...(showSkeleton?S.chipActive:{})}} onClick={() => setShowSkeleton(s=>!s)}>🦴 Esqueleto</button>
            <button style={{...S.chip,...(privacyMask?S.chipActive:{})}} onClick={() => setPrivacyMask(s=>!s)}>🔒 Privacidade</button>
            <button style={{...S.chip,...(romMode?S.chipActive:{})}} onClick={() => {setRomMode(r=>!r);setRomPoints([])}}>📏 ROM</button>
          </div>
          {romMode && (
            <div style={S.romInfo}>
              <p style={{fontSize:12,color:"#94a3b8",margin:0}}>Toque 3 pontos para medir ângulo. {romPoints.length}/3 {romPoints.length===3&&` — ${angleBetween(romPoints[0],romPoints[1],romPoints[2]).toFixed(1)}°`}</p>
              {romPoints.length>0&&<button style={S.miniBtn} onClick={()=>setRomPoints([])}>Limpar</button>}
            </div>
          )}
          {hasKp && selectedKp !== null && (
            <div style={S.kpInfo}>
              <span style={{fontSize:13,color:"#22d3ee",fontWeight:600}}>{KEYPOINT_LABELS[KEYPOINT_NAMES[selectedKp]]}</span>
              <span style={{fontSize:11,color:"#94a3b8"}}>Confiança: {(keypoints[selectedKp]?.confidence*100).toFixed(0)}% — Arraste para ajustar</span>
            </div>
          )}
          {hasKp && <p style={{fontSize:11,color:"#64748b",textAlign:"center",margin:"4px 0 0"}}>💡 Arraste os pontos para ajustar</p>}
        </div>
        {customAngles.length>0 && (
          <div style={{padding:"12px 16px"}}>
            <p style={{fontWeight:600,fontSize:13,color:"#e2e8f0",margin:"0 0 6px"}}>Ângulos Medidos (ROM)</p>
            {customAngles.map((a,i)=>(<div key={i} style={S.romResultItem}><span>Medição {i+1}</span><span style={{color:"#f59e0b",fontWeight:700}}>{a.angle.toFixed(1)}°</span></div>))}
          </div>
        )}
      </div>
    );
  }

  // ─── PHOTO RESULTS ─────────────────────────────────────────────
  if (screen === "results" && analysis) {
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("editor")}>← Editor</button>
          <h2 style={S.headerTitle}>Resultados</h2>
          <button style={S.backBtn} onClick={() => setScreen("home")}>🏠</button>
        </div>
        {analysis.score !== null && (
          <div style={S.scoreCard}>
            <div style={{...S.scoreBig,color:overallColor(analysis.score)}}>{analysis.score}</div>
            <div style={S.scoreLabel}>{overallLabel(analysis.score)}</div>
            <div style={S.scoreBar}><div style={{...S.scoreBarFill,width:`${analysis.score}%`,background:overallColor(analysis.score)}}/></div>
            <p style={{fontSize:11,color:"#94a3b8",margin:"6px 0 0"}}>Vista: {VIEW_TYPES.find(v=>v.id===viewType)?.label}</p>
          </div>
        )}
        <div style={{padding:"0 16px"}}>
          {analysis.results.map((r,i)=>(
            <div key={i} style={S.resultCard}>
              <div style={S.resultHeader}>
                <span style={{fontSize:18}}>{r.emoji}</span>
                <div style={{flex:1}}><p style={S.resultName}>{r.name}</p><p style={S.resultDesc}>{r.description}</p></div>
                <div style={{textAlign:"right",flexShrink:0}}>
                  <span style={{fontSize:18,fontWeight:800,color:r.color,display:"block"}}>{r.value}</span>
                  <span style={{fontSize:10,fontWeight:700,color:r.color,textTransform:"uppercase"}}>{r.level}</span>
                </div>
              </div>
              <div style={S.miniBar}><div style={{...S.miniBarFill,width:`${Math.min(r.deviation*3,100)}%`,background:r.color}}/></div>
            </div>
          ))}
        </div>
        <div style={{padding:"20px 16px 0"}}>
          <h3 style={{fontSize:15,fontWeight:700,color:"#e2e8f0",margin:"0 0 12px"}}>💪 Exercícios Recomendados</h3>
          {analysis.results.filter(r=>r.level!=="Normal").slice(0,3).map((r,i)=>(
            <div key={i} style={S.exerciseCard}>
              <p style={{fontWeight:600,fontSize:13,color:"#e2e8f0",margin:0}}>{r.name}</p>
              <p style={{fontSize:12,color:"#94a3b8",margin:"4px 0 0"}}>
                {r.name.includes("Cabeça")||r.name.includes("FHP") ? "Retração cervical, chin tucks, alongamento do esternocleidomastóideo." :
                 r.name.includes("Ombro") ? "Retração escapular, remada baixa, alongamento peitoral." :
                 r.name.includes("Pélvico")||r.name.includes("Tronco") ? "Ponte (glute bridge), dead bug, prancha lateral." :
                 r.name.includes("Joelho") ? "Fortalecimento do vasto medial, equilíbrio unipodal." :
                 r.name.includes("Cifose") ? "Extensão torácica, foam roller torácico." :
                 "Exercícios de mobilidade geral, yoga, pilates."}
              </p>
            </div>
          ))}
        </div>
        <p style={S.disclaimer}>⚕️ Relatório indicativo. Consulte um profissional de saúde.</p>
      </div>
    );
  }

  // ─── VIDEO PROCESSING ──────────────────────────────────────────
  if (screen === "video_processing") {
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => {setVideoProcessing(false);setScreen("home")}}>← Cancelar</button>
          <h2 style={S.headerTitle}>🎥 Processando Vídeo</h2>
          <div style={{width:60}}/>
        </div>
        <div style={{display:"flex",flexDirection:"column",alignItems:"center",justifyContent:"center",padding:"60px 20px",gap:20}}>
          <div style={S.spinner}/>
          <p style={{color:"#e2e8f0",fontSize:16,fontWeight:600,margin:0}}>Analisando postura dinâmica...</p>
          <div style={{width:"100%",maxWidth:320}}>
            <div style={{height:8,background:"rgba(100,116,139,0.2)",borderRadius:4,overflow:"hidden"}}>
              <div style={{height:"100%",background:"linear-gradient(90deg,#6366f1,#22d3ee)",borderRadius:4,transition:"width 0.3s",width:`${videoProgress}%`}}/>
            </div>
            <p style={{fontSize:13,color:"#94a3b8",textAlign:"center",margin:"10px 0 0"}}>{videoProgress}% — Detectando pose frame a frame</p>
          </div>
          <div style={{padding:"16px 20px",background:"rgba(99,102,241,0.06)",borderRadius:12,border:"1px solid rgba(99,102,241,0.12)",maxWidth:320}}>
            <p style={{fontSize:12,color:"#94a3b8",margin:0,lineHeight:1.6}}>
              🤖 O modelo MoveNet Thunder analisa cada frame do vídeo para rastrear 17 pontos anatômicos e calcular ângulos posturais ao longo do tempo.
            </p>
          </div>
        </div>
      </div>
    );
  }

  // ─── VIDEO RESULTS ─────────────────────────────────────────────
  if (screen === "video_results" && videoAnalysis) {
    const { stats, angleTimelines, totalFrames, duration } = videoAnalysis;
    const chartData = selectedVideoAngles.map(id => angleTimelines[id] || []);
    const chartColors = selectedVideoAngles.map(id => VIDEO_ANGLES.find(a => a.id === id)?.color || "#fff");
    const chartLabels = selectedVideoAngles.map(id => VIDEO_ANGLES.find(a => a.id === id)?.name || id);

    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("home")}>← Início</button>
          <h2 style={S.headerTitle}>🎥 Análise Dinâmica</h2>
          <div style={{width:60}}/>
        </div>

        {/* Video playback with overlay */}
        <div style={{...S.canvasContainer, background:"#000"}}>
          <canvas ref={videoCanvasRef} style={{display:"block",maxWidth:"100%",width: videoFrames[0]?.width || 400, height: videoFrames[0]?.height || 300}}/>
        </div>

        {/* Playback controls */}
        <div style={{padding:"12px 16px"}}>
          <div style={{display:"flex",alignItems:"center",gap:10,marginBottom:8}}>
            <button style={{...S.chip,...S.chipActive,padding:"8px 16px"}} onClick={() => setVideoPlaying(p => !p)}>
              {videoPlaying ? "⏸ Pausar" : "▶️ Reproduzir"}
            </button>
            <span style={{fontSize:12,color:"#94a3b8"}}>
              Frame {videoFrameIdx + 1}/{videoFrames.length} • {((videoFrames[videoFrameIdx]?.time) || 0).toFixed(1)}s
            </span>
          </div>
          <input type="range" min={0} max={videoFrames.length - 1} value={videoFrameIdx}
            onChange={e => { setVideoPlaying(false); setVideoFrameIdx(Number(e.target.value)); }}
            style={{width:"100%",accentColor:"#6366f1"}}
          />
        </div>

        {/* Angle selection */}
        <div style={{padding:"0 16px 12px"}}>
          <p style={{fontSize:12,fontWeight:600,color:"#94a3b8",margin:"0 0 8px"}}>Ângulos exibidos:</p>
          <div style={{display:"flex",gap:6,flexWrap:"wrap"}}>
            {VIDEO_ANGLES.map(a => {
              const active = selectedVideoAngles.includes(a.id);
              return (
                <button key={a.id} style={{...S.chip,...(active?{background:`${a.color}22`,borderColor:a.color,color:a.color}:{})}}
                  onClick={() => setSelectedVideoAngles(prev => active ? prev.filter(x=>x!==a.id) : [...prev,a.id])}>
                  {a.name}
                </button>
              );
            })}
          </div>
        </div>

        {/* Timeline chart */}
        <div style={{padding:"0 16px 16px"}}>
          <div style={{background:"rgba(15,15,30,0.6)",borderRadius:14,padding:16,border:"1px solid rgba(100,116,139,0.1)"}}>
            <p style={{fontSize:14,fontWeight:700,color:"#e2e8f0",margin:"0 0 12px"}}>📈 Variação Angular ao Longo do Tempo</p>
            <MiniChart data={chartData} colors={chartColors} labels={chartLabels} height={160} />
          </div>
        </div>

        {/* Stats cards */}
        <div style={{padding:"0 16px"}}>
          <p style={{fontSize:14,fontWeight:700,color:"#e2e8f0",margin:"0 0 12px"}}>📊 Estatísticas da Análise</p>
          <div style={{display:"flex",gap:8,marginBottom:12,flexWrap:"nowrap",overflowX:"auto",paddingBottom:4}}>
            <div style={S.statMini}>
              <span style={{fontSize:20,fontWeight:800,color:"#6366f1"}}>{totalFrames}</span>
              <span style={{fontSize:10,color:"#94a3b8"}}>Frames</span>
            </div>
            <div style={S.statMini}>
              <span style={{fontSize:20,fontWeight:800,color:"#22d3ee"}}>{duration?.toFixed(1)}s</span>
              <span style={{fontSize:10,color:"#94a3b8"}}>Duração</span>
            </div>
            <div style={S.statMini}>
              <span style={{fontSize:20,fontWeight:800,color:"#f59e0b"}}>{videoAnalysis.fps}</span>
              <span style={{fontSize:10,color:"#94a3b8"}}>FPS análise</span>
            </div>
          </div>

          {VIDEO_ANGLES.map(angleDef => {
            const s = stats[angleDef.id];
            if (!s) return null;
            const rangeSev = classifySeverity(s.range, [10, 25, 45]);
            return (
              <div key={angleDef.id} style={S.resultCard}>
                <div style={{display:"flex",alignItems:"center",gap:10,marginBottom:8}}>
                  <div style={{width:4,height:32,borderRadius:2,background:angleDef.color}}/>
                  <div>
                    <p style={{fontSize:14,fontWeight:700,color:"#e2e8f0",margin:0}}>{angleDef.name}</p>
                    <p style={{fontSize:11,color:"#94a3b8",margin:"2px 0 0"}}>
                      Variação: {rangeSev.emoji} {rangeSev.level} ({s.range.toFixed(1)}°)
                    </p>
                  </div>
                </div>
                <div style={{display:"flex",gap:16}}>
                  <div><span style={{fontSize:10,color:"#64748b",display:"block"}}>Média</span><span style={{fontSize:16,fontWeight:700,color:angleDef.color}}>{s.avg.toFixed(1)}°</span></div>
                  <div><span style={{fontSize:10,color:"#64748b",display:"block"}}>Mín</span><span style={{fontSize:16,fontWeight:700,color:"#22c55e"}}>{s.min.toFixed(1)}°</span></div>
                  <div><span style={{fontSize:10,color:"#64748b",display:"block"}}>Máx</span><span style={{fontSize:16,fontWeight:700,color:"#f87171"}}>{s.max.toFixed(1)}°</span></div>
                  <div><span style={{fontSize:10,color:"#64748b",display:"block"}}>Amplitude</span><span style={{fontSize:16,fontWeight:700,color:"#eab308"}}>{s.range.toFixed(1)}°</span></div>
                </div>
                <div style={{...S.miniBar,marginTop:10}}><div style={{...S.miniBarFill,width:`${Math.min(s.range*2,100)}%`,background:angleDef.color}}/></div>
              </div>
            );
          })}
        </div>

        {/* Interpretation */}
        <div style={{padding:"16px 16px 0"}}>
          <div style={{background:"rgba(99,102,241,0.06)",borderRadius:14,padding:16,border:"1px solid rgba(99,102,241,0.12)"}}>
            <p style={{fontSize:14,fontWeight:700,color:"#e2e8f0",margin:"0 0 8px"}}>🔍 Interpretação</p>
            {Object.entries(stats).map(([id, s]) => {
              const angleDef = VIDEO_ANGLES.find(a => a.id === id);
              if (!angleDef) return null;
              let interp = "";
              if (id === "fhp") {
                interp = s.avg < 8 ? "Cabeça bem posicionada durante o movimento." : s.avg < 18 ? "Anteriorização moderada — atenção à posição cervical." : "Anteriorização significativa — recomenda-se fortalecimento cervical.";
              } else if (id === "trunk") {
                interp = s.avg < 8 ? "Tronco bem alinhado durante o movimento." : s.avg < 15 ? "Leve inclinação do tronco durante o movimento." : "Inclinação significativa — possível fraqueza do core.";
              } else if (id === "knee") {
                interp = s.range < 15 ? "Joelho estável durante o movimento." : s.range < 40 ? "Boa amplitude de movimento do joelho." : "Grande amplitude — verificar controle motor.";
              } else if (id === "hip") {
                interp = s.range < 15 ? "Quadril estável durante o movimento." : s.range < 35 ? "Boa mobilidade do quadril." : "Grande amplitude — verificar estabilidade pélvica.";
              }
              return (
                <p key={id} style={{fontSize:12,color:"#c7d2fe",margin:"0 0 6px",lineHeight:1.5}}>
                  <span style={{color:angleDef.color,fontWeight:600}}>{angleDef.name}:</span> {interp}
                </p>
              );
            })}
          </div>
        </div>

        <p style={{...S.disclaimer,padding:"20px 16px"}}>⚕️ Análise dinâmica indicativa. Consulte um profissional para avaliação completa.</p>
      </div>
    );
  }

  // ─── HISTORY ───────────────────────────────────────────────────
  if (screen === "history") {
    return (
      <div style={S.app}>
        <div style={S.header}>
          <button style={S.backBtn} onClick={() => setScreen("home")}>← Voltar</button>
          <h2 style={S.headerTitle}>📊 Histórico</h2>
          <div style={{width:60}}/>
        </div>
        {history.length === 0 ? (
          <p style={{textAlign:"center",color:"#94a3b8",padding:40}}>Nenhuma avaliação registrada.</p>
        ) : (
          <div style={{padding:"0 16px"}}>
            {[...history].reverse().map((h,i) => (
              <div key={i} style={S.historyCard}>
                <div style={{display:"flex",justifyContent:"space-between",alignItems:"center"}}>
                  <div>
                    <p style={{fontWeight:600,fontSize:13,color:"#e2e8f0",margin:0}}>{h.date}</p>
                    <p style={{fontSize:11,color:"#94a3b8",margin:"2px 0 0"}}>{VIEW_TYPES.find(v=>v.id===h.viewType)?.label}</p>
                  </div>
                  {h.score !== null && <div style={{fontSize:28,fontWeight:900,color:overallColor(h.score)}}>{h.score}</div>}
                </div>
                <div style={{display:"flex",gap:6,flexWrap:"wrap",marginTop:8}}>
                  {h.results.slice(0,4).map((r,j)=>(<span key={j} style={{fontSize:10,padding:"3px 8px",borderRadius:6,border:`1px solid ${r.color}`,color:r.color,fontWeight:600}}>{r.emoji} {r.name.split(" ")[0]}</span>))}
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    );
  }

  return null;
}

// ─── Styles ──────────────────────────────────────────────────────
const S = {
  app: { minHeight:"100vh",background:"linear-gradient(170deg,#0f0f1a 0%,#13132b 40%,#0c1220 100%)",color:"#e2e8f0",fontFamily:"'Segoe UI',-apple-system,BlinkMacSystemFont,sans-serif",maxWidth:480,margin:"0 auto",paddingBottom:32,position:"relative" },
  homeContainer: { padding:"24px 20px" },
  logoArea: { textAlign:"center",marginBottom:24 },
  logoIcon: { display:"inline-block",marginBottom:8,filter:"drop-shadow(0 0 20px rgba(99,102,241,0.3))" },
  logoTitle: { fontSize:32,fontWeight:800,margin:0,background:"linear-gradient(135deg,#6366f1,#22d3ee)",WebkitBackgroundClip:"text",WebkitTextFillColor:"transparent" },
  logoSub: { fontSize:13,color:"#94a3b8",margin:"4px 0 0",letterSpacing:0.5 },
  tipCard: { display:"flex",gap:12,alignItems:"center",padding:"14px 16px",background:"rgba(99,102,241,0.08)",borderRadius:14,border:"1px solid rgba(99,102,241,0.15)",marginBottom:24,minHeight:52 },
  tipText: { fontSize:13,color:"#c7d2fe",margin:0,lineHeight:1.5 },
  homeActions: { display:"flex",flexDirection:"column",gap:14,marginBottom:28 },
  primaryBtn: { display:"flex",alignItems:"center",justifyContent:"center",gap:10,padding:"16px 24px",fontSize:16,fontWeight:700,color:"#fff",background:"linear-gradient(135deg,#6366f1,#4f46e5)",border:"none",borderRadius:14,cursor:"pointer",boxShadow:"0 4px 20px rgba(99,102,241,0.3)" },
  secondaryBtn: { display:"flex",alignItems:"center",justifyContent:"center",gap:8,padding:"12px 20px",fontSize:14,fontWeight:600,color:"#c7d2fe",background:"rgba(99,102,241,0.1)",border:"1px solid rgba(99,102,241,0.2)",borderRadius:12,cursor:"pointer" },
  viewSelector: { padding:"0 4px" },
  viewLabel: { fontSize:12,color:"#94a3b8",margin:"0 0 8px",fontWeight:600 },
  viewGrid: { display:"grid",gridTemplateColumns:"1fr 1fr 1fr 1fr",gap:8 },
  viewBtn: { display:"flex",flexDirection:"column",alignItems:"center",gap:3,padding:"10px 6px",background:"rgba(30,30,50,0.6)",border:"1px solid rgba(100,116,139,0.2)",borderRadius:10,color:"#94a3b8",cursor:"pointer",fontSize:16 },
  viewBtnActive: { background:"rgba(99,102,241,0.15)",borderColor:"#6366f1",color:"#c7d2fe",boxShadow:"0 0 12px rgba(99,102,241,0.2)" },
  features: { display:"flex",flexDirection:"column",gap:10,marginBottom:24 },
  featureCard: { display:"flex",gap:14,alignItems:"center",padding:"14px 16px",background:"rgba(20,20,40,0.5)",borderRadius:12,border:"1px solid rgba(100,116,139,0.1)" },
  featureTitle: { fontSize:14,fontWeight:700,color:"#e2e8f0",margin:0 },
  featureDesc: { fontSize:11,color:"#94a3b8",margin:"2px 0 0" },
  disclaimer: { fontSize:11,color:"#64748b",textAlign:"center",lineHeight:1.5,padding:"0 16px" },
  loadingBar: { display:"flex",alignItems:"center",gap:8,justifyContent:"center",marginTop:16 },
  loadingPulse: { width:8,height:8,borderRadius:"50%",background:"#6366f1",animation:"pulse 1.5s infinite" },

  header: { display:"flex",alignItems:"center",justifyContent:"space-between",padding:"14px 16px",borderBottom:"1px solid rgba(100,116,139,0.15)",background:"rgba(15,15,30,0.8)",backdropFilter:"blur(10px)",position:"sticky",top:0,zIndex:10 },
  headerTitle: { fontSize:15,fontWeight:700,margin:0,color:"#c7d2fe" },
  backBtn: { background:"none",border:"none",color:"#818cf8",fontSize:13,fontWeight:600,cursor:"pointer",padding:"4px 8px" },

  canvasContainer: { position:"relative",display:"flex",justifyContent:"center",background:"#0a0a15",minHeight:200 },
  canvas: { display:"block",touchAction:"none" },
  loadingOverlay: { position:"absolute",inset:0,display:"flex",flexDirection:"column",alignItems:"center",justifyContent:"center",background:"rgba(10,10,20,0.85)",zIndex:5 },
  spinner: { width:40,height:40,border:"3px solid rgba(99,102,241,0.2)",borderTop:"3px solid #6366f1",borderRadius:"50%",animation:"spin 0.8s linear infinite" },

  tools: { padding:"12px 16px" },
  toolRow: { display:"flex",gap:10,marginBottom:10 },
  toolBtn: { flex:1,padding:"12px 16px",fontSize:13,fontWeight:700,background:"rgba(30,30,55,0.8)",border:"1px solid rgba(99,102,241,0.25)",borderRadius:10,color:"#c7d2fe",cursor:"pointer",textAlign:"center" },
  toolBtnAccent: { background:"linear-gradient(135deg,#6366f1,#4f46e5)",border:"none",color:"#fff" },
  chip: { flex:1,padding:"8px 10px",fontSize:12,fontWeight:600,background:"rgba(20,20,40,0.6)",border:"1px solid rgba(100,116,139,0.2)",borderRadius:20,color:"#94a3b8",cursor:"pointer",textAlign:"center" },
  chipActive: { background:"rgba(99,102,241,0.15)",borderColor:"#6366f1",color:"#a5b4fc" },
  romInfo: { display:"flex",alignItems:"center",justifyContent:"space-between",padding:"8px 12px",background:"rgba(245,158,11,0.08)",border:"1px solid rgba(245,158,11,0.2)",borderRadius:8,marginBottom:8 },
  miniBtn: { background:"none",border:"1px solid rgba(245,158,11,0.3)",color:"#f59e0b",fontSize:11,padding:"3px 10px",borderRadius:6,cursor:"pointer" },
  kpInfo: { display:"flex",flexDirection:"column",gap:2,padding:"8px 12px",background:"rgba(34,211,238,0.06)",borderRadius:8,border:"1px solid rgba(34,211,238,0.15)" },
  romResultItem: { display:"flex",justifyContent:"space-between",padding:"8px 12px",background:"rgba(30,30,50,0.6)",borderRadius:8,marginBottom:6,fontSize:13,color:"#e2e8f0" },

  scoreCard: { textAlign:"center",padding:"28px 20px 20px",margin:"16px 16px 20px",borderRadius:20,background:"linear-gradient(170deg,rgba(20,20,45,0.9),rgba(15,15,35,0.9))",border:"1px solid rgba(99,102,241,0.15)",boxShadow:"0 8px 32px rgba(0,0,0,0.3)" },
  scoreBig: { fontSize:64,fontWeight:900,lineHeight:1 },
  scoreLabel: { fontSize:16,fontWeight:700,color:"#c7d2fe",marginTop:4 },
  scoreBar: { height:6,background:"rgba(100,116,139,0.2)",borderRadius:3,marginTop:14,overflow:"hidden" },
  scoreBarFill: { height:"100%",borderRadius:3,transition:"width 0.8s ease" },

  resultCard: { padding:"14px 16px",background:"rgba(20,20,40,0.5)",borderRadius:14,border:"1px solid rgba(100,116,139,0.1)",marginBottom:10 },
  resultHeader: { display:"flex",gap:12,alignItems:"flex-start" },
  resultName: { fontSize:14,fontWeight:700,color:"#e2e8f0",margin:0 },
  resultDesc: { fontSize:11,color:"#94a3b8",margin:"3px 0 0",lineHeight:1.4 },
  miniBar: { height:3,background:"rgba(100,116,139,0.15)",borderRadius:2,marginTop:10,overflow:"hidden" },
  miniBarFill: { height:"100%",borderRadius:2,transition:"width 0.6s ease" },

  exerciseCard: { padding:"12px 14px",background:"rgba(34,197,94,0.06)",borderRadius:10,border:"1px solid rgba(34,197,94,0.15)",marginBottom:8 },

  historyCard: { padding:"14px 16px",background:"rgba(20,20,40,0.5)",borderRadius:14,border:"1px solid rgba(100,116,139,0.1)",marginBottom:10 },

  statMini: { flex:1,display:"flex",flexDirection:"column",alignItems:"center",padding:"12px 8px",background:"rgba(20,20,40,0.5)",borderRadius:12,border:"1px solid rgba(100,116,139,0.1)" },
};

if (typeof document !== "undefined" && !document.getElementById("posture-ai-styles")) {
  const style = document.createElement("style");
  style.id = "posture-ai-styles";
  style.textContent = `
    @keyframes spin { to { transform: rotate(360deg) } }
    @keyframes pulse { 0%,100%{opacity:1;transform:scale(1)} 50%{opacity:0.5;transform:scale(0.8)} }
    * { box-sizing: border-box; -webkit-tap-highlight-color: transparent; }
    body { margin: 0; background: #0f0f1a; }
    button:hover { filter: brightness(1.1); }
    button:active { transform: scale(0.97); }
    input[type=range] { height: 4px; }
    input[type=range]::-webkit-slider-thumb { -webkit-appearance: none; width: 16px; height: 16px; border-radius: 50%; background: #6366f1; cursor: pointer; }
  `;
  document.head.appendChild(style);
}
