package tidbclient

import (
	"time"
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/journeymidnight/yig/error"
	. "github.com/journeymidnight/yig/meta/types"
	)

func (t *TidbClient) GetBucketNameByDomain(domain string) (bucketName string, err error) {
	domainName := ""
	sqltext := fmt.Sprintf("select * from cd where domain='%s';", domain)
	err = t.Client.QueryRow(sqltext).Scan(
		&domainName,
		&bucketName,
	)
	if err != nil && err == sql.ErrNoRows {
		err = ErrNoSuchDomain
		return
	} else if err != nil {
		return
	}
	return
}

func (t *TidbClient) SetBucketDomain(domain, bucketName string) error {
	domainName := ""
	sqltext := fmt.Sprintf("select * from cd where domain='%s';", domain)
	err = t.Client.QueryRow(sqltext).Scan(
		&domainName,
		&bucketName,
	)
	if err != nil && err == sql.ErrNoRows {
		err = ErrNoSuchDomain
		return
	} else if err != nil {
		return
	}
	return
}

func (t *TidbClient) DeleteBucketDomain(domain, bucketName string) error {
	domainName := ""
	sqltext := fmt.Sprintf("select * from cd where domain='%s';", domain)
	err = t.Client.QueryRow(sqltext).Scan(
		&domainName,
		&bucketName,
	)
	if err != nil && err == sql.ErrNoRows {
		err = ErrNoSuchDomain
		return
	} else if err != nil {
		return
	}
	return
}

func (t *TidbClient) GetDomainsByBucketName(bucketName string) (domains []string, err error) {
	sqltext := fmt.Sprintf("select * from cd where bucketname='%s';", bucketName)
	rows, err := t.Client.Query(sqltext)
	if err == sql.ErrNoRows {
		err = nil
		return
	} else if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var domain, bucket string
		err = rows.Scan(
			&domain,
			&bucket,
			)
		if err != nil {
			return
		}
		domains = append(domains, domain)
	}
	return
}