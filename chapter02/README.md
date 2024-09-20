# cURL requests

Post a recipe:

    curl --location --request POST 'http://localhost:8001/recipes' --header 'Content-Type: application/json' --data-raw '{"name": "Homemade Pizza", "tags" : ["italian", "pizza", "dinner"], "ingredients": ["1 1/2 cups (355 ml) warm water (105°F-115°F)", "1 package (2 1/4 teaspoons) of active dry yeast", "3 3/4 cups (490 g) bread flour", "feta cheese, firm mozzarella cheese, grated"], "instructions": ["Step 1.", "Step 2.", "Step 3."]}' | jq -r

Get a list of all recipes:

    curl -s --location --request GET 'http://localhost:8001/recipes' --header 'Content-Type: application/json' | jq -r

Count number of recipes:

    curl -s -X GET 'http://localhost:8001/recipes' | jq length

Change a recipe:

    curl --location --request PUT 'http://localhost:8001/recipes/cr5r02f600cd1vtqqvng' --header 'Content-Type: application/json' --data-raw '{"id": "cr5r02f600cd1vtqqvng", "name": "Shrimp Scampi Pizza", "tags" : ["italian", "pizza", "dinner"], "ingredients": ["1 1/2 cups (355 ml) warm water (105°F-115°F)", "1 package (2 1/4 teaspoons) of active dry yeast", "3 3/4 cups (490 g) bread flour", "firm mozzarella cheese, grated, shrimps"], "instructions": ["Step 1.", "Step 2.", "Step 3."]}' | jq -r

Delete a recipe:
    
    curl -v -X DELETE http://localhost:8001/recipes/c0283p3d0cvuglq85log | jq -r

Search tags:

    curl -v -X GET http://localhost:8001/recipes/search?tag=italian