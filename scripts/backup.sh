#!/bin/bash

REMOTE_HOST="..."
REMOTE_PB_DATA="/srv/www/pocketbase"
LOCAL_PB_DATA="/var/www/pocketbase"

# --------------------------------------------------------------------
# Simple shellscript to backup the remote PocketBase (data.db + storage)
# to a local directory.
#
# NB! Update the above variables according to your system.
#
# To add it as a cron task (crontab -e):
# ```sh
# # run on every 12h
# 0 */12 * * * /path/to/app/backup.sh
# ```
# --------------------------------------------------------------------

echo ""
echo "---------------------------------------------------------------"
echo "-> SQLite database backup"
echo "---------------------------------------------------------------"

# create a db backup file
ssh -T root@$REMOTE_HOST << EOL
cd $REMOTE_PB_DATA
sqlite3 data.db '.backup data.db.backup'
exit
EOL

if [ $? -ne 0 ]; then
    exit 1
fi

# rsync the db backup file
rsync -av -e ssh root@$REMOTE_HOST:$REMOTE_PB_DATA/data.db.backup $LOCAL_PB_DATA/data.db

if [ $? -ne 0 ]; then
    exit 1
fi

echo ""
echo "---------------------------------------------------------------"
echo "-> Storage files backup"
echo "---------------------------------------------------------------"

# rsync the storage files
rsync -av -e ssh root@$REMOTE_HOST:$REMOTE_PB_DATA/storage/ $LOCAL_PB_DATA/storage/

if [ $? -ne 0 ]; then
    exit 1
fi

echo ""
echo "==============================================================="
echo "Backup completed!"
