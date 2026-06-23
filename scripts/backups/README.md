# Backups Plenya (banco + uploads)

Sistema de backup de produção e o espelho local. **Fonte versionada** dos scripts que
rodam na VPS e nesta máquina de dev — se algo se perder, restaurar a partir daqui.

> ⚠️ Estes scripts vivem em runtime FORA do repo (na VPS em `/usr/local/bin/`, no dev em
> `/home/user/backups/`). As cópias aqui são a referência canônica. Ao editar, atualize
> os dois lugares.

## Visão geral (3 camadas para PROD)

1. **VPS gera + rotaciona** (cron root):
   - `plenya-db-backup.sh` → `db_plenya_*.dump` (`pg_dump -Fc`), **03:00**, mantém **14**.
   - `plenya-uploads-backup.sh` → `uploads_*.tar.gz` (mídia WhatsApp + `patient-docs`), **03:30**, mantém **7**.
   - Saída: `/home/deploy/.plenya-vps-secrets/backups/` na VPS.
2. **Dev espelha** (`pull-prod-backups.sh`, cron local `15 9 * * *`, best-effort — WSL não fica 24/7):
   - `rsync --delete` da pasta da VPS → `/home/user/backups/prod/vps/`.
3. **Guardar pra sempre:** mover o dump de `prod/vps/` para `prod/` (fora do espelho com `--delete`).

## Local ÚNICO dos backups nesta máquina
```
/home/user/backups/
├── dev/        # dumps do banco DEV (manuais)
├── prod/       # dumps do banco PROD (manuais/históricos)
│   └── vps/    # espelho automático da VPS — NÃO editar à mão
```
**Regra:** sempre salvar backup do Plenya aqui. Não criar pasta de backup solta.

## Instalar na VPS (se precisar refazer)
```bash
scp scripts/backups/plenya-db-backup.sh plenya:/tmp/ && \
  ssh plenya "sudo install -m755 /tmp/plenya-db-backup.sh /usr/local/bin/"
scp scripts/backups/plenya-uploads-backup.sh plenya:/tmp/ && \
  ssh plenya "sudo install -m755 /tmp/plenya-uploads-backup.sh /usr/local/bin/"
# cron root:
ssh plenya "( sudo crontab -l 2>/dev/null; printf '%s\n%s\n' '0 3 * * * /usr/local/bin/plenya-db-backup.sh' '30 3 * * * /usr/local/bin/plenya-uploads-backup.sh' ) | sort -u | sudo crontab -"
```
UUID do app api (path do bind dos uploads) e nome do container Postgres estão nos scripts;
se mudarem, ver memória `plenya_vps_emr`.

## Restaurar
```bash
# Banco (custom -Fc):
pg_restore -U plenya_user -d plenya_db --clean --if-exists db_plenya_<ts>.dump
# Uploads:
tar -xzf uploads_<ts>.tar.gz -C <uploadsRoot>   # ex: /app/uploads dentro do container
```

## Pendência conhecida
Tudo hoje é disco local (VPS + esta máquina). Falta **off-site real** (S3/R2/outro host).
Ver `docs/emr/plano-persistencia-uploads-vps.md` e memória `plenya_backups_local_canonico`.
