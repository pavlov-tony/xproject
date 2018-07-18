CREATE OR REPLACE FUNCTION xproject.sel_last_report()
RETURNS TABLE (id integer, name text, bucket text, time_created timestamp without time zone, account_id integer) AS $$
BEGIN
	RETURN QUERY
	SELECT *
	FROM xproject.gcp_csv_files AS reps
	WHERE reps.time_created = (SELECT MAX(max_reps.time_created) FROM xproject.gcp_csv_reports AS max_reps);
END; $$
LANGUAGE plpgsql;

--CREATE OR REPLACE FUNCTION xproject.sel_last_report()
--RETURNS SETOF xproject.gcp_csv_reports AS
--$$
--BEGIN
--	SELECT * FROM xproject.gcp_csv_reports;
--END;
--$$ LANGUAGE plpgsql;
