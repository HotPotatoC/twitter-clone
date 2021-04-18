-- Sources:
-- https://www.postgresql.org/message-id/20050810133157.GA46247@winnie.fuhr.org (plpgsql function by Michael Fuhr)
-- https://wiki.postgresql.org/wiki/Count_estimate
-- https://www.citusdata.com/blog/2016/10/12/count-performance/ (Good article)
BEGIN;
CREATE FUNCTION count_estimate_rows (query text)
    RETURNS integer
    AS $$
DECLARE
    rec record;
    ROWS integer;
BEGIN
    FOR rec IN EXECUTE 'EXPLAIN ' || query LOOP
        -- Parses the EXPLAIN query to get the estimate count of the rows
        ROWS := SUBSTRING(rec. "QUERY PLAN" FROM ' rows=([[:digit:]]+)');
        EXIT
        WHEN ROWS IS NOT NULL;
    END LOOP;
    RETURN ROWS;
    END;
$$
LANGUAGE plpgsql
VOLATILE STRICT;
COMMIT;

