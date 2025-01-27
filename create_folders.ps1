$folders = @(
    "01_basics/variables_and_types",
    "01_basics/functions",
    "01_basics/control_structures",
    "01_basics/arrays_and_slices",
    "01_basics/maps",
    "01_basics/structs_and_methods",
    "02_intermediate/pointers",
    "02_intermediate/interfaces",
    "02_intermediate/concurrency/goroutines",
    "02_intermediate/concurrency/channels",
    "02_intermediate/concurrency/sync",
    "02_intermediate/error_handling",
    "03_advanced/file_handling",
    "03_advanced/testing",
    "03_advanced/generics",
    "03_advanced/reflection",
    "04_building_cli_tools/cobra",
    "04_building_cli_tools/flag",
    "04_building_cli_tools/custom_tools",
    "05_web_development/http_server",
    "05_web_development/rest_api",
    "05_web_development/gorilla_mux",
    "06_database/sql",
    "06_database/orm",
    "06_database/nosql",
    "07_deployment/building_binaries",
    "07_deployment/docker",
    "07_deployment/ci_cd",
    "07_deployment/versioning"
)

# Create all directories
foreach ($folder in $folders) {
    New-Item -ItemType Directory -Force -Path $folder
}
