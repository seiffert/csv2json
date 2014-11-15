CSV2JSON
=======================

This little tool converts a CSV file with 2 lines into a JSON string. My use Case: publishing PHPLoc results to Fluentd.

Development:
-----------------------

Use make to create the pseiffert/csv2json container.

	make

Run it:
-----------------------

    docker run -it --rm -v $(pwd):/opt/workspace pseiffert/phploc --log-csv=phploc.csv $PHP_PROJECT_DIRECTORY 1>/dev/null && \
        cat phploc.csv | \
        docker run -i --rm pseiffert/csv2json