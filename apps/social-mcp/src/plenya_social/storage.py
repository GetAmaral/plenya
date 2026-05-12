import json
import os
import tempfile
from pathlib import Path

from cryptography.fernet import Fernet

from . import config


def _fernet() -> Fernet:
    return Fernet(config.ENCRYPTION_KEY)


def _path(name: str) -> Path:
    return config.TOKENS_DIR / f"{name}.enc"


def save(name: str, data: dict) -> None:
    payload = _fernet().encrypt(json.dumps(data, separators=(",", ":")).encode())
    target = _path(name)
    fd, tmp = tempfile.mkstemp(prefix=f"{name}.", suffix=".tmp", dir=target.parent)
    try:
        with os.fdopen(fd, "wb") as f:
            f.write(payload)
        os.replace(tmp, target)
        os.chmod(target, 0o600)
    except Exception:
        if os.path.exists(tmp):
            os.unlink(tmp)
        raise


def load(name: str) -> dict | None:
    target = _path(name)
    if not target.exists():
        return None
    return json.loads(_fernet().decrypt(target.read_bytes()).decode())


def exists(name: str) -> bool:
    return _path(name).exists()
