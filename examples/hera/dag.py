from hera.workflows import DAG, WorkflowTemplate, script

@script()
def echo(message: str):
    print(message)

with WorkflowTemplate(
    generate_name="dag-example",
    entrypoint="diamond",
) as w:
    with DAG(name="diamond"):
        A = echo(name="A", arguments={"message": "A"})
        B = echo(name="B", arguments={"message": "B"})
        C = echo(name="C", arguments={"message": "C"})
        D = echo(name="D", arguments={"message": "D"})
        A >> [B, C] >> D

w.to_file('.')
