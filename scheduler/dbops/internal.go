package dbops

import "log"

func ReadVideoDeletionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("select video_id from video_del_rec LIMIT ?")
	var ids []string
	if err != nil {
		return ids, err
	}
	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDleetionRecord error:%v", err)
		return ids, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}

func DelVideoDeletionRecode(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE from video_del_rec where video_id = ?")

	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletionRecodrd error:%v", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
